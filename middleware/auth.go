package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"RESTApi/utils"

)

func Authenticate(context *gin.Context){
	token:= context.Request.Header.Get("Authorization")
	if token == ""{
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "User is not authorized"})
		context.Abort()
		return
	} 
	
	userID, err := utils.ValidateToken(token)
	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "User not authorized"})
		context.Abort()
		return
	}

	context.Set("userID", userID)
	context.Next()
	
}