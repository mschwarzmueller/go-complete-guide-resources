package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /events/5
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
}
