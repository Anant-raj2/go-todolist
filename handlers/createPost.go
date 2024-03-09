package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
