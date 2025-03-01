package main

import (
	"database/sql"

	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rvif/cli-todo-app/cmd"
)

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		color.Red("🔴 DB_URL is not found in .env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		color.Red("🔴 Error connecting to the database: %v", err)
		return
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		color.Red("🔴 Error pinging the database: %v", err)
		return
	}
	// color.Green("🟩 Connected to the database")
	cmd.Execute(conn)
}
