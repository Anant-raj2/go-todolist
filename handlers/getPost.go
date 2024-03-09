package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Post",
	})
}

func GetAllPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Post",
	})
}
