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
	authenticate.POST("/events/:id/register", registrationForAnEvent)
	authenticate.DELETE("/events/:id/register",cancellationForAnEvent)


	server.Handle("POST", "/signup", signUp)
	server.Handle("POST", "/login", logIn)
	server.Handle("GET", "/events", getEvents)
	server.Handle("GET", "/events/:id", getAnEvent)
	//server.Handle("POST", "/events/:id/register", registrationForAnEvent)
	//sqlserver.Handle("DELETE", "/events/:id/register", cancellationForAnEvent)
	
}
