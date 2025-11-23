package routes

import (
	"RESTApi/models"
	"net/http"
	"RESTApi/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.Users

	err:= context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message" : "User signed up successfully"})
}

func logIn(context *gin.Context) {
	var user models.Users

	err:= context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	err = user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "User unauthorised"})
		return
	}

	token, err := utils.TokenGeneration(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Invalid credentials"})
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Login Successful"})

	context.JSON(http.StatusOK, gin.H{"token": token })



	//context.JSON
}