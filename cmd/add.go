package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/rvif/cli-todo-app/internal/database"
)

var addCmd = &cobra.Command{
	Use:   "add [task_name]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		newId := NewShortUUID()
		istTime := getISTTime()

		queries := database.New(db)

		_, err := queries.CreateTask(cmd.Context(), database.CreateTaskParams{
			ID:          string(newId),
			Name:        taskName,
			IsCompleted: false,
			CreatedAt:   istTime,
			UpdatedAt:   istTime,
		})

		if err != nil {
			log.Fatalf("Error creating task: %v", err)
		}

		fmt.Printf("Task added: [%s] %s\n", newId, taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
