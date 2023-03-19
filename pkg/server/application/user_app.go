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
	CreateUser(*entities.User) (string, error)
	GetUserByID(string) (*entities.User, error)
	GetUserByUsername(string) (*entities.User, error)
	GetUserByCredentials(string, string) (*entities.User, error)
	SearchByUsername(string) ([]*entities.User, error)
	UpdateUser(*entities.User) (*entities.User, error)
	DeleteUser(string) error
	UpdateUserOnlineStatus(string, int) error
}

func (u *UserApp) CreateUser(user *entities.User) (string, error) {
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

func (u *UserApp) SearchByUsername(username string) ([]*entities.User, error) {
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
