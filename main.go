package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"robertgleason.ca/rest/models"
)

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = 1
	event.UserID = 1
	c.JSON(http.StatusCreated, gin.H{"status": "success", "event": event})
}

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

}
