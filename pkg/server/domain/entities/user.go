package entities

import (
	"errors"
	"project-eighteen/pkg/constants"
)

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string
	Password string
	Online   string
	SocketID string
}

type UserDetailsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDetailsResponse struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Online   string `json:"online"`
}

type Users []User

func (users Users) PublicUsers() []interface{} {
	res := make([]interface{}, len(users))
	for i, user := range users {
		res[i] = user.PublicUser()
	}
	return res
}

func (user *User) PublicUser() interface{} {
	return &UserDetailsResponse{
		UserID:   user.ID,
		Username: user.Username,
		Online:   user.Online,
	}
}

func (user *User) Validate() error {
	if user.Username == "" {
		return errors.New(constants.UsernameCantBeEmpty)
	}
	if user.Password == "" {
		return errors.New(constants.PasswordCantBeEmpty)
	}
	if len(user.Username) < 3 || len(user.Username) > 20 {
		return errors.New(constants.UsernameLenError)
	}

	return nil
}

func (user *UserDetailsRequest) Validate() error {
	if user.Username == "" {
		return errors.New(constants.UsernameCantBeEmpty)
	}
	if user.Password == "" {
		return errors.New(constants.PasswordCantBeEmpty)
	}
	if len(user.Username) < 3 || len(user.Username) > 20 {
		return errors.New(constants.UsernameLenError)
	}

	return nil
}