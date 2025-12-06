package routes

import (
	"RESTApi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registrationForAnEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10,64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID passed"})
		return
	}
	_, err = models.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "No search event found"})
		return
	}

	userid:= context.GetInt64("userID")

	err = models.RegisterForEvent(userid, eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "User successfully registered for event"})
}
