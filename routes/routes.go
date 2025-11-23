package routes

import (
	"RESTApi/middleware"

	"github.com/gin-gonic/gin"
)

func RouterInitialisation(server *gin.Engine) {
	authenticate:= server.Group("/")
	authenticate.Use(middleware.Authenticate)

	
	authenticate.POST( "/events", CreateEvents)
	authenticate.DELETE( "/events/:id", deleteAnEvent)
	authenticate.PUT("/events/:id", updateAnEvent)


	server.Handle("POST", "/signup", signUp)
	server.Handle("POST", "/login", logIn)
	server.Handle("GET", "/events", getEvents)
	server.Handle("GET", "/events/:id", getAnEvent)
	

}
