package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Anant-raj2/todolist/data"
)

func CreatePost(c *gin.Context) {
	var post data.Post
	if err := c.BindJSON(&post); err != nil {
		panic(err)
	}
	data.CreatePost(post)
	c.Status(http.StatusCreated)
}
