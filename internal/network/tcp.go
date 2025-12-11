package network

import (
	"fmt"
	"goback/internal/network/client"
	"goback/internal/network/server"
	"io"
	"log"
	"net"
	"os"
)

func StartServer(port string) {

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	defer listener.Close()

	log.Printf("Listening on port: %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}
		server.SetConnection(conn)
		go server.ListenToClient()
		go server.SendToClient()
	}
}

func StartClient(address string, port string, keyChan <-chan string) {

	log.Printf("Connecting to: %s:%s", address, port)
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal("Error connecting:", err)
	}
	defer conn.Close()

	log.Println("Connection Succeded")

	client.SetConnection(conn)
	go client.ListenToServer()

	for key := range keyChan {
		fmt.Fprintf(conn, "%s", key)
	}
	io.Copy(os.Stdout, conn)
}
