package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"project/internal/handler"
)

var (
	host     = "localhost"
	connType = "tcp"
	port     = "8989"
)

func main() {
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) >= 2 {
		fmt.Printf("[USAGE]: ./TCPChat $port\n")
		return
	}
	listen, err := net.Listen(connType, host+":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Listening on the port :%v\n", port)

	defer func() {
		log.Printf("End listen tcp port: %s\n", port)
		listen.Close()
	}()

	if err = handler.Run(listen); err != nil {
		log.Fatalln(err)
	}
}
