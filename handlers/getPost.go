package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Anant-raj2/todolist/data"
)

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post := data.GetPostById(id)
	if post == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, post)
	}
}

func GetAllPosts(c *gin.Context) {
	var posts []data.Post = data.GetAllPosts()
	if posts == nil || len(posts) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, posts)
	}
}
