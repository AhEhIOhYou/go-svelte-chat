package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"project-eighteen/pkg/redisrepo"

	"github.com/gin-gonic/gin"
)

type userReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Client   string `json:"client"`
}

type response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data.omitempty"`
	Total   int         `json:"total.omitempty"`
}

func registerHandler(c *gin.Context) {
	var u userReq
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := register(&u)
	c.JSON(http.StatusOK, res)
}

func loginHandler(c *gin.Context) {
	var u userReq
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := login(&u)
	c.JSON(http.StatusOK, res)
}

func verifyContactHandler(c *gin.Context) {
	username := c.Param("username")

	res := verifyContact(username)
	c.JSON(http.StatusOK, res)
}

func chatHistoryHandler(c *gin.Context) {
	username1, err := c.GetQuery("username1")
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username1 not found"})
		return
	}

	username2, err := c.GetQuery("username2")
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username2 not found"})
		return
	}

	fromTS, toTS := "0", "+inf"

	if c.Query("from-ts") != "" && c.Query("to-ts") != "" {
		fromTS = c.Query("from-ts")
		toTS = c.Query("to-ts")
	}

	res := chatHistory(username1, username2, fromTS, toTS)
	c.JSON(http.StatusOK, res)
}

func contactListHandler(c *gin.Context) {
	username := c.Param("username")

	res := contactList(username)
	c.JSON(http.StatusOK, res)
}

func register(u *userReq) *response {
	res := &response{Status: true}

	status := redisrepo.IsUserExist(u.Username)
	if status {
		res.Status = false
		res.Message = "Username already taken"
		return res
	}

	err := redisrepo.RegisterNewUser(u.Username, u.Password)
	if err != nil {
		res.Status = false
		res.Message = "Failed to register new user, try again later"
		return res
	}

	return res
}

func login(u *userReq) *response {
	res := &response{Status: true}

	err := redisrepo.IsUserAuthentic(u.Username, u.Password)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return res
	}

	return res
}

func verifyContact(username string) *response {
	res := &response{Status: true}

	status := redisrepo.IsUserExist(username)
	if !status {
		res.Status = false
		res.Message = "Username not found"
		return res
	}

	return res
}

func chatHistory(username1, username2, fromTS, toTS string) *response {
	res := &response{Status: true}

	fmt.Println(username1, username2)

	if !redisrepo.IsUserExist(username1) || !redisrepo.IsUserExist(username2) {
		res.Status = false
		res.Message = "Username not found"
		return res
	}

	chats, err := redisrepo.FetchChatBetween(username1, username2, fromTS, toTS)
	if err != nil {
		res.Status = false
		log.Println("Fetch chat history error: ", err)
		res.Message = "Unable to fetch chat history, try again later"
		return res
	}

	res.Data = chats
	res.Total = len(chats)

	return res
}

func contactList(username string) *response {
	res := &response{Status: true}

	contacts, err := redisrepo.FetchContactList(username)
	if err != nil {
		res.Status = false
		log.Println("Fetch contact list error: ", err)
		res.Message = "Unable to fetch contact list, try again later"
		return res
	}

	res.Data = contacts
	res.Total = len(contacts)

	return res
}
