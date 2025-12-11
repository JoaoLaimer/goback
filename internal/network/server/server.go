package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var connection net.Conn

func SendToClient() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		text := scanner.Text()
		fmt.Fprintf(connection, text+"\n")
		fmt.Print("\r>> ")

	}
}
func ListenToClient() {
	defer connection.Close()
	scanner := bufio.NewScanner(connection)
	fmt.Print("\r>> ")
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf(": %s", text)
	}

}

func SetConnection(conn net.Conn) {
	connection = conn
}
