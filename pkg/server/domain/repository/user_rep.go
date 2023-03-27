package repository

import (
	"project-eighteen/pkg/server/domain/entities"
)

type UserRepository interface {
	CreateUser(*entities.UserDetailsRequest) (string, error)
	GetUserByID(string) (*entities.User, error)
	GetUserByUsername(string) (*entities.User, error)
	GetUserByCredentials(string, string) (*entities.User, error)
	SearchByUsername(string) ([]*entities.UserDetailsResponse, error)
	UpdateUser(*entities.User) (*entities.User, error)
	DeleteUser(string) error
	UpdateUserOnlineStatus(string, int) error
}