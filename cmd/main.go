package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"project/internal/handler"
)

var (
	connType = "tcp"

	host string
	port string
)

func main() {
	flag.StringVar(&port, "port", "8989", "Port on which to run")
	flag.StringVar(&host, "host", "localhost", "The address where it will be launched")
	flag.Parse()

	listen, err := net.Listen(connType, host+":"+port)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Listening on the port: %v\n", port)

	// defer func() {
	// 	log.Printf("End listen tcp port: %s\n", port)
	// 	listen.Close()
	// }()

	if err = handler.Run(listen); err != nil {
		log.Fatalln(err)
	}
}
