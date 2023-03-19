package interfaces

import (
	"fmt"
	"net/http"
	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/server/application"
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/infrastructure/security"
	"project-eighteen/pkg/server/infrastructure/utils"

	"github.com/gin-gonic/gin"
)

type Users struct {
	userApp application.UserAppInterface
}

func NewUsersHandler(userApp application.UserAppInterface) *Users {
	return &Users{userApp: userApp}
}

func (u *Users) Register(ctx *gin.Context) {
	var user entities.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	validateErr := user.Validate()
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, validateErr),
		})
		return
	}

	userObjID, regErr := u.userApp.CreateUser(&user)
	if regErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, regErr),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
		"userID":  userObjID,
	})
}

func (u *Users) Login(ctx *gin.Context) {
	var user entities.UserDetailsRequest

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	validiteErr := user.Validate()
	if validiteErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	userDetails, err := u.userApp.GetUserByUsername(user.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
	}

	if userDetails == (&entities.User{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.AccountDoesNotExist,
		})
		return
	}

	if err := security.VerifyPassword(userDetails.Password, user.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.PasswordIncorrect,
		})
		return
	}

	u.userApp.UpdateUserOnlineStatus(userDetails.ID, 1)

	ctx.JSON(http.StatusOK, gin.H{
		"message": &entities.UserDetailsResponse{
			UserID:   userDetails.ID,
			Username: userDetails.Username,
			Online:   userDetails.Online,
		},
	})
}

func (u *Users) IsUsernameAvailable(ctx *gin.Context) {
	username := ctx.Param("username")

	if !utils.CheckAlphaNumeric(username) {
		ctx.JSON(http.StatusOK, gin.H{
			"avaliable": false,
			"message":   constants.UsernameInvalid,
		})
	} else if len(username) < 3 || len(username) > 20 {
		ctx.JSON(http.StatusOK, gin.H{
			"avaliable": false,
			"message":   constants.UsernameLenError,
		})
	} else {
		userDetails, err := u.userApp.GetUserByUsername(username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"avaliable": false,
				"message":   constants.Failed,
			})
			return
		}
		if userDetails == (&entities.User{}) {
			ctx.JSON(http.StatusOK, gin.H{
				"avaliable": true,
				"message":   constants.UsernameIsAvailable,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"avliable": false,
				"message":  constants.UsernameIsAlreadyTaken,
			})
		}
	}
}

func (u *Users) SearchUser(ctx *gin.Context) {
	username := ctx.Query("username")

	if !utils.CheckAlphaNumeric(username) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.UsernameInvalid,
		})
	}

	users, err := u.userApp.SearchByUsername(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
		"users":   users,
	})
}

func (u *Users) Logout(ctx *gin.Context) {
	userID := ctx.Param("userID")

	u.userApp.UpdateUserOnlineStatus(userID, 0)

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
	})
}