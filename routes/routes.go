package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes (server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.AuthMiddleware)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", createUser)
	server.POST("/login", loginUser)
}
