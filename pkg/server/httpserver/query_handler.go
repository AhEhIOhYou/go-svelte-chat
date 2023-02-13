package httpserver

import (
	"context"
	"errors"
	"os"
	"time"

	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/database"
	"project-eighteen/pkg/server/structs"
	"project-eighteen/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserOnlineStatus(userID string, status string) error {
	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil
	}

	collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, queryError := collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"$set": bson.M{"online": status}})
	defer cancel()

	if queryError != nil {
		return errors.New(constants.Failed)
	}
	return nil
}

func GetUserByUsername(username string) structs.UserDetailsType {
	var userDetails structs.UserDetailsType

	collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = collection.FindOne(ctx, bson.M{"username": username}).Decode(&userDetails)

	defer cancel()

	return userDetails
}

func GetUserByUserID(userID string) structs.UserDetailsType {
	var userDetails structs.UserDetailsType

	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return userDetails
	}

	collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&userDetails)

	defer cancel()

	return userDetails
}

func IsUsernameAvaliableQH(username string) bool {
	userDetails := GetUserByUsername(username)
	if userDetails == (structs.UserDetailsType{}) {
		return true
	}
	return false
}

func LoginQH(user structs.UserDetailsRequestPayloadType) (structs.UserDetailsResponsePayloadType, error) {
	if user.Username == "" {
		return structs.UserDetailsResponsePayloadType{}, errors.New(constants.UsernameCantBeEmpty)
	} else if user.Password == "" {
		return structs.UserDetailsResponsePayloadType{}, errors.New(constants.PasswordCantBeEmpty)
	} else {
		userDetails := GetUserByUsername(user.Username)

		if userDetails == (structs.UserDetailsType{}) {
			return structs.UserDetailsResponsePayloadType{}, errors.New(constants.AccountDoesNotExist)
		}

		if err := utils.VerifyPassword(userDetails.Password, user.Password); err != nil {
			return structs.UserDetailsResponsePayloadType{}, errors.New(constants.PasswordIncorrect)
		}

		if err := UpdateUserOnlineStatus(userDetails.ID, "Y"); err != nil {
			return structs.UserDetailsResponsePayloadType{}, errors.New(constants.PasswordIncorrect)
		}

		return structs.UserDetailsResponsePayloadType{
			UserID:   userDetails.ID,
			Username: userDetails.Username,
		}, nil
	}
}

func RegisterQH(user structs.UserDetailsRequestPayloadType) (string, error) {
	if user.Username == "" {
		return "", errors.New(constants.UsernameCantBeEmpty)
	} else if user.Password == "" {
		return "", errors.New(constants.PasswordCantBeEmpty)
	} else {
		newPassword, err := utils.Hash(user.Password)
		if err != nil {
			return "", errors.New(constants.Failed)
		}
		collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		regRes, regErr := collection.InsertOne(ctx, bson.M{
			"username": user.Username,
			"password": newPassword,
			"online":   "N",
		})
		defer cancel()

		regObjectID, regObjErr := regRes.InsertedID.(primitive.ObjectID)

		if onlineStatusErr := UpdateUserOnlineStatus(regObjectID.Hex(), "Y"); onlineStatusErr != nil {
			return "", errors.New(constants.Failed)
		}

		if regErr != nil || !regObjErr {
			return "", errors.New(constants.Failed)
		}

		return regObjectID.Hex(), nil
	}
}

func GetAllOnlineUsers(userID string) []structs.UserDetailsResponsePayloadType {
	var onlineUsers []structs.UserDetailsResponsePayloadType

	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return onlineUsers
	}

	collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, queryErr := collection.Find(ctx, bson.M{
		"online": "Y",
		"_id":    bson.M{"$ne": docID},
	})
	defer cancel()

	if queryErr != nil {
		return onlineUsers
	}

	for cursor.Next(context.TODO()) {
		var user structs.UserDetailsType
		err := cursor.Decode(&user)
		if err == nil {
			onlineUsers = append(onlineUsers, structs.UserDetailsResponsePayloadType{
				UserID:   user.ID,
				Username: user.Username,
				Online:   user.Online,
			})
		}
	}

	return onlineUsers
}

func GetAllUsers(userID string) []structs.UserDetailsResponsePayloadType {
	var onlineUsers []structs.UserDetailsResponsePayloadType

	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return onlineUsers
	}

	collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, queryErr := collection.Find(ctx, bson.M{
		"_id": bson.M{"$ne": docID},
	})
	defer cancel()

	if queryErr != nil {
		return onlineUsers
	}

	for cursor.Next(context.TODO()) {
		var user structs.UserDetailsType
		err := cursor.Decode(&user)
		if err == nil {
			onlineUsers = append(onlineUsers, structs.UserDetailsResponsePayloadType{
				UserID:   user.ID,
				Username: user.Username,
				Online:   user.Online,
			})
		}
	}

	return onlineUsers
}

func StoreNewChatMessage(message structs.MessagePayloadType) bool {
	collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, regErr := collection.InsertOne(ctx, bson.M{
		"from":    message.From,
		"to":      message.To,
		"message": message.Message,
	})
	defer cancel()

	if regErr != nil {
		return false
	}

	return true
}

func GetChatMessages(from, to string) []structs.ChatType {
	var chatMessages []structs.ChatType

	collection := database.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, queryErr := collection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"$and": []bson.M{{"to": to}, {"from": from}}},
			{"$and": []bson.M{{"to": from}, {"from": to}}},
		},
	})
	defer cancel()

	if queryErr != nil {
		return chatMessages
	}

	for cursor.Next(context.TODO()) {
		var chat structs.ChatType
		err := cursor.Decode(&chat)
		if err == nil {
			chatMessages = append(chatMessages, structs.ChatType{
				ID:      chat.ID,
				From:    chat.From,
				To:      chat.To,
				Message: chat.Message,
			})
		}
	}

	return chatMessages
}
