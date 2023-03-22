package persistence

import (
	"context"
	"fmt"
	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

var _ repository.UserRepository = &UserRepo{}

func NewUserRepo(collection *mongo.Collection) *UserRepo {
	return &UserRepo{collection: collection}
}

func (r *UserRepo) CreateUser(user *entities.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	regRes, regErr := r.collection.InsertOne(ctx, bson.M{
		"username": user.Username,
		"password": user.Password,
		"online":   "N",
	})
	defer cancel()

	_, regObjectErr := regRes.InsertedID.(primitive.ObjectID)

	if regErr != nil || !regObjectErr {
		return "", fmt.Errorf(constants.DatabaseError, regErr)
	}

	return regRes.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *UserRepo) GetUserByID(userID string) (*entities.User, error) {
	var user entities.User

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = r.collection.FindOne(ctx, bson.M{"_id": userDocID}).Decode(&user)

	defer cancel()

	return &user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	defer cancel()

	return &user, nil
}

func (r *UserRepo) GetUserByCredentials(username string, password string) (*entities.User, error) {
	var user entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = r.collection.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user)

	defer cancel()

	return &user, nil
}

func (r *UserRepo) SearchByUsername(username string) ([]*entities.User, error) {
	var users []*entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"username": bson.M{"$regex": username}})
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user entities.User
		cursor.Decode(&user)
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepo) UpdateUser(user *entities.User) (*entities.User, error) {
	var userUp entities.User = *user

	docID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	var status string

	if user.Online == "Y" {
		status = "N"
	} else {
		status = "Y"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, queryError := r.collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"$set": bson.M{"online": status}})

	if queryError != nil {
		return nil, fmt.Errorf(constants.DatabaseError, queryError)
	}

	return &userUp, nil
}

func (r *UserRepo) DeleteUser(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryError := r.collection.DeleteOne(ctx, bson.M{"_id": userDocID})

	if queryError != nil {
		return fmt.Errorf(constants.DatabaseError, queryError)
	}

	return nil
}

func (r *UserRepo) UpdateUserOnlineStatus(userID string, status int) error {
	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, queryError := r.collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"$set": bson.M{"online": status}})

	if queryError != nil {
		return fmt.Errorf(constants.DatabaseError, queryError)
	}

	return nil
}