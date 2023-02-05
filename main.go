package main

import (
	"flag"
	"fmt"
	"log"
	"project-eighteen/pkg/httpserver"
	"project-eighteen/pkg/websocket"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()

	if *server == "http" {
		fmt.Println("Starting HTTP Server")
		httpserver.StartServer()
	} else if *server == "websocket" {
		fmt.Println("Starting Websocket Server")
		websocket.StartServer()
	} else {
		fmt.Println("Please provide the server type")
	}
}
