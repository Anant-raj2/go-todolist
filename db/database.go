package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	var connectionString string = os.Getenv("CONNECT_STRING")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database connected successfully!")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database pinged successfully!")
	}

	defer db.Close()
}
