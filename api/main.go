package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)

	router.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Ping is working"})
}
