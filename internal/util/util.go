package util

import (
	"errors"
	"net"
	"os"
	"sync"
)

func Logo() (string, error) {
	text, err := os.ReadFile("../pkg/logo.txt")
	if err != nil {
		return "", err
	}
	return string(text), nil
}

func CheckConnection(conn net.Conn, name string, mu *sync.Mutex, Clients map[net.Conn]string) error {
	mu.Lock()
	defer mu.Unlock()
	if len(name) == 0 {
		return errors.New("The name is empty!")
	}
	for _, nameInClients := range Clients {
		if nameInClients == name {
			return errors.New("This name is already!")
		}
	}
	if len(Clients) > 9 {
		return errors.New("There are 10 clients in the chat room.")
	}
	for _, value := range name {
		if !(value >= 'a' && value <= 'z') && !(value >= 'A' && value <= 'Z') {
			return errors.New("The name have not a valid value! Type only letters.")
		}
	}
	return nil
}
