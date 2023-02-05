package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-eighteen/model"
	"project-eighteen/pkg/redisrepo"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Username string
}

type Message struct {
	Type string     `json:"type"`
	User string     `json:"user,omitempty"`
	Chat model.Chat `json:"chat,omitempty"`
}

var clients = make(map[*Client]bool)
var broadcast = make(chan *model.Chat)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r.URL.Query())

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error in upgrading connection", err)
		return
	}

	client := &Client{Conn: conn}
	clients[client] = true
	fmt.Println("clients", len(clients), clients, conn.RemoteAddr().String())

	reciver(client)

	fmt.Println("Client disconnected", client.Username, conn.RemoteAddr().String())
	delete(clients, client)
}

func reciver(client *Client) {
	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error in reading message from client", err)
			return
		}

		var message Message

		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println("Error in unmarshalling message from client", err)
			return
		}

		if message.Type == "bootup" {
			client.Username = message.User
			fmt.Println("Client connected", client.Username)
		} else {
			fmt.Println("Received message from client", message.Type, message.Chat)
			c := message.Chat
			c.Timestamp = time.Now().Unix()

			id, err := redisrepo.CreateChat(&c)
			if err != nil {
				log.Println("Error in creating chat", err)
				return
			}

			c.ID = id

			broadcast <- &c
		}
	}
}

func sender() {
	for {
		message := <-broadcast
		fmt.Println("Broadcasting message", message)
		for client := range clients {
			fmt.Println("Sending message, username", client.Username, "from", message.From, "to", message.To)

			if client.Username == message.From || client.Username == message.To {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					log.Println("Error in writing message to client", err)
					client.Conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/ws", serveWebSocket)
}

func StartServer() {
	redisClient := redisrepo.InitialiseRedis()
	defer redisClient.Close()

	go sender()
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
