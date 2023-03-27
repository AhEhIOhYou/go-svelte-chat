package entities

import (
	"errors"
	"project-eighteen/pkg/server/constants"
)

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Online   int    `json:"online" bson:"online"`
	SocketID string `json:"socketId" bson:"socketId"`
}

type UserDetailsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDetailsResponse struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Online   int    `json:"online"`
}

type Users []User

func (users Users) PublicUsers() []interface{} {
	res := make([]interface{}, len(users))
	for i, user := range users {
		res[i] = user.PublicUser()
	}
	return res
}

func (user *User) PublicUser() UserDetailsResponse {
	return UserDetailsResponse{
		ID:   user.ID,
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
