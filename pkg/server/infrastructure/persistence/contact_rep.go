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

type ContactRepo struct {
	collection *mongo.Collection
}

var _ repository.ContactRepository = &ContactRepo{}

func NewContactRepo(collection *mongo.Collection) *ContactRepo {
	return &ContactRepo{
		collection: collection,
	}
}

func (r *ContactRepo) CreateContact(contact *entities.Contact) (*entities.Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	regRes, regErr := r.collection.InsertOne(ctx, bson.M{
		"user_id": contact.UserID,
		"contact": contact.ContactID,
	})
	defer cancel()

	_, regObjectErr := regRes.InsertedID.(primitive.ObjectID)

	if regErr != nil || !regObjectErr {
		return nil, fmt.Errorf(constants.DatabaseError, regErr)
	}

	contact.ID = regRes.InsertedID.(primitive.ObjectID).Hex()

	return contact, nil
}

func (r *ContactRepo) DeleteContact(userID, contactID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	contactDocID, err := primitive.ObjectIDFromHex(contactID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryError := r.collection.DeleteOne(ctx, bson.M{"user_id": userDocID, "contact": contactDocID})

	if queryError != nil {
		return fmt.Errorf(constants.DatabaseError, queryError)
	}

	return nil
}

func (r *ContactRepo) DeleteContactsByUserID(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf(constants.DatabaseError, err)
	}

	_, queryError := r.collection.DeleteMany(ctx, bson.M{"user_id": userDocID})

	if queryError != nil {
		return fmt.Errorf(constants.DatabaseError, queryError)
	}

	return nil
}

func (r *ContactRepo) GetListOfContactsByUserID(userID string) ([]*entities.Contact, error) {
	var contacts []*entities.Contact

	userDocID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userDocID})
	if err != nil {
		return nil, fmt.Errorf(constants.DatabaseError, err)
	}

	for cursor.Next(ctx) {
		var contact entities.Contact
		cursor.Decode(&contact)
		contacts = append(contacts, &contact)
	}

	return contacts, nil
}
