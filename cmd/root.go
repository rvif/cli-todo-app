package cmd

import (
	"database/sql"
	"encoding/json"
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

// Temp json solution for storing tasks
var tasks []Task

const taskFile = "tasks.json"

// global postgress connection
var db *sql.DB

// Load tasks from JSON file (or create if missing)
func loadTasks() {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		os.WriteFile(taskFile, []byte("[]"), 0644) // 0644: read and write the file or directory and other users can only read it.
		return
	}

	file, err := os.ReadFile(taskFile)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

// Save to JSON file
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "   ")
	// log.Println(string(data))
	os.WriteFile(taskFile, data, 0644)
}

var rootCmd = &cobra.Command{
	Use:   "cli-todo-app",
	Short: "A simple CLI todo app",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		loadTasks() // Load tasks before running any command
	},
}

func Execute(conn *sql.DB) {
	db = conn
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
