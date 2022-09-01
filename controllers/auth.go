package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/users/utils"
)

type RegisterInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"email" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := utils.ValidateEmail(input.Email)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": input})
}
