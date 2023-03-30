package application

import (
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/domain/repository"
)

type UserApp struct {
	userRep repository.UserRepository
}

var _ UserAppInterface = &UserApp{}

type UserAppInterface interface {
	// Create a new user in the database and return the user's ID
	CreateUser(*entities.UserDetailsRequest) (string, error)
	// Get a user by ID
	GetUserByID(string) (*entities.User, error)
	// Get a user by username
	GetUserByUsername(string) (*entities.User, error)
	// Get a user by username and password
	GetUserByCredentials(string, string) (*entities.User, error)
	// Search for users by username
	SearchByUsername(string) ([]*entities.UserDetailsResponse, error)
	// Update a user
	UpdateUser(*entities.User) (*entities.User, error)
	// Delete a user
	DeleteUser(string) error
	// Update a user's online status
	UpdateUserOnlineStatus(string, int) error
}

func (u *UserApp) CreateUser(user *entities.UserDetailsRequest) (string, error) {
	return u.userRep.CreateUser(user)
}

func (u *UserApp) GetUserByID(id string) (*entities.User, error) {
	return u.userRep.GetUserByID(id)
}

func (u *UserApp) GetUserByUsername(username string) (*entities.User, error) {
	return u.userRep.GetUserByUsername(username)
}

func (u *UserApp) GetUserByCredentials(username, password string) (*entities.User, error) {
	return u.userRep.GetUserByCredentials(username, password)
}

func (u *UserApp) SearchByUsername(username string) ([]*entities.UserDetailsResponse, error) {
	return u.userRep.SearchByUsername(username)
}

func (u *UserApp) UpdateUser(user *entities.User) (*entities.User, error) {
	return u.userRep.UpdateUser(user)
}

func (u *UserApp) DeleteUser(id string) error {
	return u.userRep.DeleteUser(id)
}

func (u *UserApp) UpdateUserOnlineStatus(id string, status int) error {
	return u.userRep.UpdateUserOnlineStatus(id, status)
}
