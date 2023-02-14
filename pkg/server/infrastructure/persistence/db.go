package persistence

import (
	"context"
	"fmt"
	"log"
	"os"
	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/server/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repos struct {
	ChatRepository    repository.ChatRepository
	ContactRepository repository.ContactRepository
	MessageRepository repository.MessageRepository
	UserRepository    repository.UserRepository
	db                *mongo.Client
}

func NewRepo(dbURI string) (*Repos, error) {
	log.Println(constants.DatabaseConnectionStart)
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseConnectionError, err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseConnectionError, err)
	}
	log.Println(constants.DatabaseConnectionSuccess)

	userCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	chatCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection("chats")
	contactCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection("contacts")
	messageCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")

	return &Repos{
		ChatRepository:    NewChatRepo(chatCollection),
		ContactRepository: NewContactRepo(contactCollection),
		MessageRepository: NewMessageRepo(messageCollection),
		UserRepository:    NewUserRepo(userCollection),
		db:                client,
	}, nil
}
