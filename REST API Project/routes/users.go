package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	
	err := context.ShouldBindJSON(&user)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		fmt.Println(err)
		return
	}
	
	err = user.Save()
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	
	err := context.ShouldBindJSON(&user)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		fmt.Println(err)
		return
	}
	
	err = user.ValidateCredentials()
	
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	
	// JWT Token Logic
	token, err := utils.GenerateToken(user.Email, user.ID)
	
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

