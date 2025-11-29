package main

import (
	"flag"
	"fmt"
	"goback/internal/keyboard"
	"goback/internal/network"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}
	mode := os.Args[1]

	clientCmd := flag.NewFlagSet("client", flag.ExitOnError)
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)

	switch mode {
	case "client":

		clientHost := clientCmd.String("h", "127.0.0.1", "Host IP")

		clientCmd.Parse(os.Args[2:])
		args := clientCmd.Args()

		if len(args) < 1 {
			fmt.Println("Usage: ./goback client -h {ip} {port}")
			os.Exit(1)
		}

		port := args[0]

		keyChan := make(chan string, 100)

		go network.StartClient(*clientHost, port, keyChan)
		//fmt.Printf("Dummy Client: %s %s\n", *clientHost, port)
		//os.Exit(0)

		keyboard.Setup(keyChan)

	case "server":

		serverPort := serverCmd.String("p", "8080", "Listening Port")

		serverCmd.Parse(os.Args[2:])
		args := serverCmd.Args()

		if *serverPort == "" || len(args) != 0 {
			fmt.Println("Usage: ./goback server -p {port}")
			os.Exit(1)
		}

		network.StartServer(*serverPort)

	default:
		log.Fatal("Not Working as Intended")
	}

}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("\t./goback client -h {ip} {port}")
	fmt.Println("\t./goback server -p {port}")

}
