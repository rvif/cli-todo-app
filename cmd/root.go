package cmd

import (
	"database/sql"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type ShortUUID string

func NewShortUUID() ShortUUID {
	return ShortUUID(uuid.New().String()[:8])
}

func getISTTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc)
}

type Task struct {
	ID          ShortUUID `json:"id"`
	Name        string    `json:"name"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// global postgress connection
var db *sql.DB

var rootCmd = &cobra.Command{
	Use:   "cli-todo-app",
	Short: "A simple CLI todo app",
}

func Execute(conn *sql.DB) {
	db = conn
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
