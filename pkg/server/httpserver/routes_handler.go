package httpserver

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/server/structs"
)

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func IsUsernameAvailable(ctx *gin.Context) {
	username := ctx.Param("username")

	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	if !IsAlphaNumeric(username) {
		ctx.JSON(http.StatusOK, gin.H{
			"available": false,
			"message": constants.UsernameInvalid,
		})
		return
	} else if len(username) < 3 || len(username) > 20 {
		ctx.JSON(http.StatusOK, gin.H{
			"available": false,
			"message": constants.UsernameLenError,
		})
		return
	} else {
		isUsernameAvailable := IsUsernameAvaliableQH(username)
		if isUsernameAvailable {
			ctx.JSON(http.StatusOK, gin.H{
				"available": true,
				"message": constants.UsernameIsAvailable,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"available": false,
				"message": constants.UsernameIsAlreadyTaken,
			})
		}
	}
}

func Registration(ctx *gin.Context) {
	var user structs.UserDetailsRequestPayloadType
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	}

	if user.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.UsernameCantBeEmpty,
		})
		return
	} else if user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.PasswordCantBeEmpty,
		})
		return
	} else {
		userObjID, regErr := RegisterQH(user)
		if regErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": constants.Failed,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"userId":  userObjID,
				"message": constants.RegistrationSuccessful,
			})
		}
	}
}

func Login(ctx *gin.Context) {
	var user structs.UserDetailsRequestPayloadType
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	}

	if user.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.UsernameCantBeEmpty,
		})
		return
	} else if user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.PasswordCantBeEmpty,
		})
		return
	} else {
		UserDetails, loginErr := LoginQH(user)
		if loginErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": constants.Failed,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    UserDetails,
				"message": constants.LoginSuccessful,
			})
		}
	}
}

func UserSessionCheck(ctx *gin.Context) {
	userID := ctx.Query("userID")

	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	if !IsAlphaNumeric(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	} else {
		userDetails := GetUserByUserID(userID)
		if userDetails == (structs.UserDetailsType{}) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": constants.NotLoggedIn,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    userDetails,
				"message": constants.YouAreLoggedIn,
			})
		}
	}
}

func GetMessagesHandler(ctx *gin.Context) {
	from := ctx.Param("from")
	to := ctx.Param("to")

	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	if !IsAlphaNumeric(from) || !IsAlphaNumeric(to) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	} else {
		messages := GetChatMessages(from, to)
		ctx.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	}
}
