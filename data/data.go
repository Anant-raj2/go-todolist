package data

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	// ExpiresAt string `json:"expires_at"`
}

var Posts = []Post{
	{ID: "1", Title: "Blue Train"},
	{ID: "2", Title: "Jeru"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown"},
}
