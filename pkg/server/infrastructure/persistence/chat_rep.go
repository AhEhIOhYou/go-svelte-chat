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
)

type ChatRepo struct {
	collection *mongo.Collection
}

var _ repository.ChatRepository = &ChatRepo{}

func NewChatRepo(collection *mongo.Collection) *ChatRepo {
	return &ChatRepo{collection: collection}
}

func (r *ChatRepo) CreateChat(chat *entities.Chat) (*entities.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	participantsDocID := make([]primitive.ObjectID, len(chat.ParticipantsID))

	for i, id := range chat.ParticipantsID {
		participantDocID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf(constants.DatabaseError, err)
		}
		
		participantsDocID[i] = participantDocID
	}

	if len(participantsDocID) > 2 {
		chat.Type = constants.ChatTypeGroup
	} else {
		chat.Type = constants.ChatTypeDialog
	}

	createRes, regErr := r.collection.InsertOne(ctx, bson.M{
		"participants": participantsDocID,
		"type":         chat.Type,
		"name":         chat.Name,
	})

	_, regObjectErr := createRes.InsertedID.(primitive.ObjectID)

	if regErr != nil || !regObjectErr {
		return nil, fmt.Errorf(constants.DatabaseError, regErr)
	}

	chat.ID = createRes.InsertedID.(primitive.ObjectID).Hex()

	return chat, nil
}

func (r *ChatRepo) CheckExsistDialog(participantsID []string) (*entities.Chat, error) {
	var chat entities.Chat

	participantsDocID := make([]primitive.ObjectID, len(participantsID))

	for i, id := range participantsID {
		participantDocID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf(constants.DatabaseError, err)
		}

		participantsDocID[i] = participantDocID
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = r.collection.FindOne(ctx, bson.M{
		"type": constants.ChatTypeDialog,
		"participants": bson.M{
			"$all":  participantsDocID,
			"$size": len(participantsDocID),
		},
	}).Decode(&chat)

	defer cancel()

	return &chat, nil
}

func (r *ChatRepo) GetChatByID(chatID string) (*entities.Chat, error) {
	var chat entities.Chat

	chatDocID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = r.collection.FindOne(ctx, bson.M{"_id": chatDocID}).Decode(&chat)

	defer cancel()

	return &chat, nil
}

func (r *ChatRepo) GetChatsByParticipantID(userID string) ([]*entities.Chat, error) {
	var chats []*entities.Chat

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"participants": userDocID})
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	for cursor.Next(ctx) {
		var chat entities.Chat
		cursor.Decode(&chat)
		chats = append(chats, &chat)
	}

	return chats, nil
}

func (r *ChatRepo) UpdateChat(chat *entities.Chat) (*entities.Chat, error) {
	var chatUp entities.Chat

	docID, err := primitive.ObjectIDFromHex(chat.ID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, queryErr := r.collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"$set": bson.M{
		"name": chat.Name,
	}})
	defer cancel()

	if queryErr != nil {
		return nil, fmt.Errorf(constants.DatabaseError, queryErr)
	}

	return &chatUp, nil
}

func (r *ChatRepo) AddParticipantToChat(chatID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chatDocID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryErr := r.collection.UpdateOne(ctx, bson.M{"_id": chatDocID}, bson.M{"$push": bson.M{
		"participants": userDocID,
	}})

	if queryErr != nil {
		return fmt.Errorf(constants.DatabaseError, queryErr)
	}

	return nil
}

func (r *ChatRepo) DeleteParticipantFromChat(chatID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chatDocID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryErr := r.collection.UpdateOne(ctx, bson.M{"_id": chatDocID}, bson.M{"$pull": bson.M{
		"participants": userDocID,
	}})

	if queryErr != nil {
		return fmt.Errorf(constants.DatabaseError, queryErr)
	}

	return nil
}

func (r *ChatRepo) DeleteChat(chatID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chatDocID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryErr := r.collection.DeleteOne(ctx, bson.M{"_id": chatDocID})

	if queryErr != nil {
		return fmt.Errorf(constants.DatabaseError, queryErr)
	}

	return nil
}
