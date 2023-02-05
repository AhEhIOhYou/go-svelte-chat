package main

import (
	"flag"
	"fmt"
	"log"

	"project-eighteen/pkg/database"
	server "project-eighteen/pkg/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	opt := flag.String("server", "", "http,websocket")
	flag.Parse()

	switch *opt {
	case "http":
		fmt.Println("Starting http server")
		database.DBConnect()
		server.StartHttpServer()
	case "websocket":
		fmt.Println("Starting websocket server")
		database.DBConnect()
		server.StartWebSocketServer()
	case "all":
		fmt.Println("Starting http and websocket servers")
		database.DBConnect()
		go server.StartWebSocketServer()
		server.StartHttpServer()
	default:
		fmt.Println("Please specify the server to start")
	}
}
