package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func listenToServer(conn net.Conn) {
	fmt.Printf("listening\n")
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("SERVER: %s", text)
	}

}
