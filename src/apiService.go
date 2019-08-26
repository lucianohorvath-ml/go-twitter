package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"github.com/lucianohorvath-ml/go-twitter/src/service"
	"net/http"
)

var writer service.TweetWriter
var tweetManager *service.TweetManager

func main() {
	setUp()
	router := gin.Default()
	registerRoutes(router)

	_ = router.Run()
}

func setUp() {
	writer = service.NewMemoryTweetWriter()
	tweetManager = service.NewTweetManager(writer)

}

func registerRoutes(router *gin.Engine) {
	router.GET("/hola/:nombre", saludar)
	router.POST("/register", registerUser)
	router.GET("/timeline", showTimeLine)
	router.GET("/users", showUsers)

}

func saludar(c *gin.Context) {
	nombre := c.Param("nombre")
	c.String(http.StatusOK, "Hola "+nombre)
}

func showTimeLine(c *gin.Context) {
	//c.JSON()
}

func showUsers(c *gin.Context) {
	users := tweetManager.UserManager.GetUsers()
	var usersString string
	for _, u := range users {
		usersString += u.Nombre
	}
	c.String(http.StatusOK, usersString)
}

func registerUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		tweetManager.UserManager.RegisterUser(&user)
		c.String(http.StatusOK, "Usuario registrado")
		return
	}
}
