package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Event Routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	
	// Grouping Routes & Adding Middleware
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	
	// Registration Routes
	authenticated.POST("/events/:id/register", registerForEvents)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	
	// User Routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}

