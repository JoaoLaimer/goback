package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func sendToClient(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		text := scanner.Text()
		fmt.Fprintf(conn, text+"\n")
		fmt.Print("\r>> ")

	}
}
func listenToClient(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	fmt.Print("\r>> ")
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf(": %s", text)
	}

}
