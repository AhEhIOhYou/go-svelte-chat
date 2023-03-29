package interfaces

import (
	"fmt"
	"net/http"
	"project-eighteen/pkg/server/application"
	"project-eighteen/pkg/server/constants"
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/infrastructure/utils"

	"github.com/gin-gonic/gin"
)

type Contacts struct {
	contactApp application.ContactAppInterface
}

func NewContactsHandler(contactApp application.ContactAppInterface) *Contacts {
	return &Contacts{
		contactApp: contactApp,
	}
}

func (c *Contacts) AddContact(ctx *gin.Context) {
	var contact entities.Contact
	err := ctx.BindJSON(&contact)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	validateErr := contact.Validate()
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, validateErr),
		})
		return
	}

	alreadyExsist, err := c.contactApp.IsContact(contact.UserID, contact.ContactID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	if alreadyExsist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, constants.ContactAlreadyExists),
		})
		return
	}

	newContact, addErr := c.contactApp.CreateContact(&contact)
	if addErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, addErr),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
		"contact": newContact,
	})
}

func (c *Contacts) GetContacts(ctx *gin.Context) {
	userID := ctx.Param("user-id")
	if !utils.CheckAlphaNumeric(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, constants.AccountDoesNotExist),
		})
		return
	}

	contacts, err := c.contactApp.GetListOfContactsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  constants.Successful,
		"contacts": contacts,
	})
}

func (c *Contacts) DeleteContact(ctx *gin.Context) {
	var contact entities.Contact

	err := ctx.BindJSON(&contact)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	err = c.contactApp.DeleteContact(contact.UserID, contact.ContactID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
	})
}

func (c *Contacts) ClearAllUserContacts(ctx *gin.Context) {
	userID := ctx.Param("userID")
	if !utils.CheckAlphaNumeric(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, constants.AccountDoesNotExist),
		})
		return
	}

	err := c.contactApp.DeleteContactsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
	})
}
