package network

import (
	"bufio"
	"fmt"
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

		go handleConnection(conn)
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

	go func() {
		for key := range keyChan {
			fmt.Fprintf(conn, "%s", key)
		}
	}()
	io.Copy(os.Stdout, conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("%s", text)

	}

}
