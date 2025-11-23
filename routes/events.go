package routes

import (
	"RESTApi/models"
	"fmt"

	"strconv"

	//"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	allEvents, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "unable to get all events"})
	}
	context.JSON(http.StatusOK, gin.H{"Events retrieved successfully": allEvents})

}

func getAnEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID passed"})
		return
	}
	getOneEvent, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "No Search event present"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Event retrieved successfully": getOneEvent})
}

func deleteAnEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID passed"})
		return
	}
	
	getOneEvent, err := models.GetEvent(eventID)
	if int64(getOneEvent.UserID) !=  context.GetInt64("userID"){
		context.JSON(http.StatusUnauthorized, gin.H{"Message" : "User did not create event"})
		context.Abort()
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "No Search event found"})
		return
	}
	err = models.DeleteEvent(eventID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"Message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Event deleted successfully": getOneEvent})
}

func updateAnEvent(context *gin.Context) {
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
	

	var event models.Events
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}
	err = event.UpdateEvent(eventID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"Message": err})
		return
	}
	updatedEvents, err := models.GetEvent(eventID)
	if err!= nil{
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "No search event found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Event updated successfully": updatedEvents})
	
}

func CreateEvents(context *gin.Context) {
	var event models.Events

	userid:= context.GetInt64("userID")
	event.UserID = int(userid)

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Created event": event})

	
}
