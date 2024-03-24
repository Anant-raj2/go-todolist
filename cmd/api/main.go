package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Anant-raj2/todolist/db"
	"github.com/Anant-raj2/todolist/handlers"
)

func main() {
	db.Connect()
	r := gin.Default()

	r.GET("/posts/:id", handlers.GetPost)
	r.GET("/posts", handlers.GetAllPosts)
	r.POST("/create-post", handlers.CreatePost)
	r.GET("/delete-post/:id", handlers.DeletePost)

	r.Run("localhost:8000")
}
