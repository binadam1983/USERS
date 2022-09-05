package controllers

import (
	"net/http"

	"github.com/users/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	utils "github.com/users/utils"
)

type RegisterInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

var token string

func Register(c *gin.Context) {

	var input RegisterInput

	/* 	err := c.ShouldBindJSON(&input)
	   	if err != nil {
	   		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	   		return
	   	} */
	input.Email = c.PostForm("email")
	input.Password = c.PostForm("password")

	valid, err := utils.ValidateEmail(input.Email)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	user.Email = input.Email
	user.Password = input.Password

	err = user.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"Success": input.Email})
	token, err := user.AuthenticateUser()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Authentication error": err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 180, "/", "localhost", true, false)

	c.HTML(http.StatusOK, "welcome.html", gin.H{"title": input.Email})
	log.Info(http.StatusOK, gin.H{"Token": token})
}

func Login(c *gin.Context) {

	var input LoginInput

	input.Email = c.PostForm("email")
	input.Password = c.PostForm("password")

	/* 	err := c.ShouldBindJSON(&input)
	   	if err != nil {
	   		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	   		return
	   	} */

	var user model.User
	user.Email = input.Email
	user.Password = input.Password

	token, err := user.AuthenticateUser()
	if err != nil || err == http.ErrNoCookie || token == "" {
		c.JSON(http.StatusNotFound, gin.H{"Authentication error": err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 180, "/", "localhost", true, false)
	c.HTML(http.StatusOK, "welcome.html", gin.H{"title": input.Email})
	log.Info(http.StatusOK, gin.H{"Token": token})
	log.Info(input.Email)
}

func Logout(c *gin.Context) {

	_, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", "", -1, "/", "localhost", true, false)
	c.HTML(http.StatusOK, "logging-out.html", gin.H{"title": "Logged out"})

}

func GetUsers(c *gin.Context) {

	users, err := model.GetUsersFromDB()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error ": err.Error()})
	}

	//c.JSON(http.StatusFound, users)

	c.HTML(http.StatusFound, "users.html", gin.H{
		"title": "Users",
		"users": users})
}

func Homepage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Homepage"})
}
