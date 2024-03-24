package data

import (
	"database/sql"
	"time"

	"github.com/Anant-raj2/todolist/db"
)

func GetAllPosts() []Post {
	var db *sql.DB = db.Connect()
	var posts []Post = []Post{}

	results, err := db.Query("SELECT * FROM todostest")
	if err != nil {
		panic(err)
	}

	for results.Next() {
		var post Post
		err = results.Scan(&post.ID, &post.Title, &post.CreatedAt)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	defer db.Close()
	return posts
}

func GetPostById(param string) *Post {
	var db *sql.DB = db.Connect()
	defer db.Close()
	results, err := db.Query("SELECT * FROM todos WHERE id='" + param + "'")
	if err != nil {
		panic(err)
	}
	var post *Post = &Post{}
	if results.Next() {
		err = results.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err)
		}
	} else {
		return nil
	}
	return post
}

func CreatePost(post Post) {
	var db *sql.DB = db.Connect()
	// var id string = uuid.NewString()
	idNew := "2"
	var createdAt string = time.Now().Format("2006-01-02 15:04:05")
	defer db.Close()
	_, err := db.Query(
		"INSERT INTO todostest VALUES ('" + idNew + "', '" + post.Title + "', '" + createdAt + "')",
	)
	if err != nil {
		panic(err)
	}
}

func DeletePost(param string) {
	var db *sql.DB = db.Connect()
	defer db.Close()

	_, err := db.Query("DELETE FROM todostest WHERE id=" + param + "")
	if err != nil {
		panic(err)
	}
}
