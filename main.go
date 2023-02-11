package main

import (
	"log"

	"project-eighteen/pkg/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	server.Start()
}
