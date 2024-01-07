package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Epiq Dev",
	})
}

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8080")

}
