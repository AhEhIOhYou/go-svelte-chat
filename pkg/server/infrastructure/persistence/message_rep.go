package persistence

import (
	"context"
	"fmt"
	"project-eighteen/pkg/server/constants"
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageRepo struct {
	collection *mongo.Collection
}

var _ repository.MessageRepository = &MessageRepo{}

func NewMessageRepo(collection *mongo.Collection) *MessageRepo {
	return &MessageRepo{
		collection: collection,
	}
}

func (r *MessageRepo) StoreMessage(message *entities.Message) (*entities.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chatDocID, chatErr := primitive.ObjectIDFromHex(message.ChatID)
	fromDocID, fromErr := primitive.ObjectIDFromHex(message.FromID)
	if chatErr != nil || fromErr != nil {
		return nil, fmt.Errorf(constants.DatabaseError, chatErr)
	}

	createRes, regErr := r.collection.InsertOne(ctx, bson.M{
		"chat_id":    chatDocID,
		"from_id":    fromDocID,
		"from_name":  message.FromName,
		"message":    message.Message,
		"created_at": message.CreatedAt,
	})

	_, regObjectErr := createRes.InsertedID.(primitive.ObjectID)

	if regErr != nil || !regObjectErr {
		return nil, fmt.Errorf(constants.DatabaseError, regErr)
	}

	message.ID = createRes.InsertedID.(primitive.ObjectID).Hex()

	return message, nil
}

func (r *MessageRepo) GetMessagesByChatID(chatID string, limit int64, offset int64) ([]*entities.Message, error) {
	var messages []*entities.Message

	chatDocID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := r.collection.Find(ctx, bson.M{"chat_id": chatDocID}, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	})

	defer cancel()

	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	for cursor.Next(ctx) {
		var message entities.Message
		cursor.Decode(&message)
		messages = append(messages, &message)
	}

	return messages, nil
}

func (r *MessageRepo) GetMessageByID(messageID string) (*entities.Message, error) {
	var message entities.Message

	userDocID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = r.collection.FindOne(ctx, bson.M{"_id": userDocID}).Decode(&message)

	defer cancel()

	return &message, nil
}

func (r *MessageRepo) UpdateMessage(message *entities.Message) (*entities.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	docID, err := primitive.ObjectIDFromHex(message.ID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"$set": bson.M{
		"message": message.Message,
	}})

	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	return message, nil
}

func (r *MessageRepo) DeleteMessage(messageID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	messageDocID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryError := r.collection.DeleteOne(ctx, bson.M{"_id": messageDocID})

	if queryError != nil {
		return fmt.Errorf(constants.DatabaseError, queryError)
	}

	return nil
}

func (r *MessageRepo) DeleteMessagesByChatID(chatID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chatDocID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryError := r.collection.DeleteMany(ctx, bson.M{"chat_id": chatDocID})

	if queryError != nil {
		return fmt.Errorf(constants.DatabaseError, queryError)
	}

	return nil
}
