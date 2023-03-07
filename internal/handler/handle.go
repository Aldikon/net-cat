package handler

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"project/internal/util"
	"project/model"
)

var (
	MessageChannel = make(chan model.Message)
	Clients        = make(map[net.Conn]string)
	mu             sync.Mutex
)

func Run(l net.Listener) error {
	history, err := os.Create("history.txt")
	if err != nil {
		return err
	}

	go SendMessegeAll()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error in connect: %v", err)
			continue
		}
		go Handle(conn, history)
	}
}

func Handle(conn net.Conn, history *os.File) {
	defer conn.Close()
	// WRITE LOGO
	fmt.Fprintln(conn, "Welcome to TCP-Chat!")
	logo, err := util.Logo()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(conn, logo)

	// CREATE VAR
	scan := bufio.NewScanner(conn)
	var name string

	// INPUT NAME
	for {
		fmt.Fprint(conn, "[ENTER YOUR NAME]:")
		if scan.Scan() {
			name = scan.Text()
		}
		if err = util.CheckConnection(conn, name, &mu, Clients); err != nil {
			fmt.Fprintln(conn, fmt.Sprint(err)+" Plese try again.")
			continue
		}
		mu.Lock()
		Clients[conn] = name
		mu.Unlock()
		break
	}

	// PRINT IN CONNECTION HISTORY
	fileData, err := os.ReadFile(history.Name())
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(conn, string(fileData))

	fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-1-2 15:4:5"), name)

	MessageChannel <- model.NewMessage(name, fmt.Sprintf("%s has joined our chat...", name))

	// INPUT TEXT
	for scan.Scan() {
		// log.Printf("Connected %v", name)
		fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-1-2 15:4:5"), name)
		if len(strings.TrimSpace(scan.Text())) == 0 {
			continue
		}
		text := strings.TrimSpace(scan.Text())
		MessageChannel <- model.NewMessage(name, fmt.Sprintf("[%s][%s]:%s", time.Now().Format("2006-1-2 15:4:5"), name, text))
		history.WriteString(fmt.Sprintf("[%s][%s]:%s\n", time.Now().Format("2006-1-2 15:4:5"), name, text))
	}
	MessageChannel <- model.NewMessage(name, fmt.Sprintf("%s has left our chat...", name))
	mu.Lock()
	delete(Clients, conn)
	mu.Unlock()
}

func SendMessegeAll() {
	for {
		msg := <-MessageChannel
		mu.Lock()
		for conn, name := range Clients {
			if name == msg.Name {
				continue
			}
			timeNow := time.Now().Format("2006-1-2 15:4:5")
			fmt.Fprintf(conn, "\n%s\n", msg.Text)
			fmt.Fprintf(conn, "[%s][%s]:", timeNow, name)
		}
		mu.Unlock()
	}
}
