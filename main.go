package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/rvif/cli-todo-app/cmd"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Println("DB_URL is not found in .env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	fmt.Println("Connected to the database")
	cmd.Execute(conn)
}
