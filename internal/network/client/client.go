package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var command string
var connection net.Conn

func ListenToServer() {

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		command := scanner.Text()
		log.Printf("SERVER: %s", command)
		ClientProcessCommand(command)
	}

}

func ClientProcessCommand(cmd string) {
	switch cmd {
	case "quit":
		closeConnection()
		os.Exit(0)
	default:
	}
}

func SetConnection(conn net.Conn) {
	connection = conn
}

func closeConnection() {
	err := connection.Close()
	if err != nil {
		fmt.Errorf("Error closing connection: %v", err)
		os.Exit(1)
	}
}
