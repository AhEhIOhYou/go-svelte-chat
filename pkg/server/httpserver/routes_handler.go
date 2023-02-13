package httpserver

import (
	"net/http"
	"regexp"

	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/server/structs"

	"github.com/gin-gonic/gin"
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
			"message":   constants.UsernameInvalid,
		})
		return
	} else if len(username) < 3 || len(username) > 20 {
		ctx.JSON(http.StatusOK, gin.H{
			"available": false,
			"message":   constants.UsernameLenError,
		})
		return
	} else {
		isUsernameAvailable := IsUsernameAvaliableQH(username)
		if isUsernameAvailable {
			ctx.JSON(http.StatusOK, gin.H{
				"available": true,
				"message":   constants.UsernameIsAvailable,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"available": false,
				"message":   constants.UsernameIsAlreadyTaken,
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
	from := ctx.Param("userFromId")
	to := ctx.Param("userToId")

	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	if !IsAlphaNumeric(from) || !IsAlphaNumeric(to) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	} else {
		messages := GetChatMessages(from, to)
		fromData := GetUserByUserID(from)
		toData := GetUserByUserID(to)

		ctx.JSON(http.StatusOK, gin.H{
			"details": gin.H{
				"fromName": fromData.Username,
				"fromId":   fromData.ID,
				"toName":   toData.Username,
				"toId":     toData.ID,
			},
			"messages": messages,
		})
	}
}

func SearchUser(ctx *gin.Context) {
	username := ctx.Query("username")

	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	if !IsAlphaNumeric(username) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	} else {
		users := SearchUserByUsername(username)
		ctx.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func AddContact(ctx *gin.Context) {
	var contact structs.ContactRequestPayloadType
	err := ctx.BindJSON(&contact)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	}

	if contact.UserID == "" || contact.ContactID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	} else {
		succes := AddContactQH(contact)
		if succes != true {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": constants.Failed,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": constants.ContactAdded,
			})
		}
	}
}

func GetContacts(ctx *gin.Context) {
	userID := ctx.Query("userID")

	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	if !IsAlphaNumeric(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.Failed,
		})
		return
	} else {
		contacts := GetContactList(userID)
		ctx.JSON(http.StatusOK, gin.H{
			"contacts": contacts,
		})
	}
}