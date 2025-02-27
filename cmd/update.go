package cmd

import (
	"fmt"
	"log"

	"github.com/rvif/cli-todo-app/internal/database"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [task_id]",
	Short: "Update a todo by its ID",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		task, err := database.New(db).GetTaskByID(cmd.Context(), id)
		if err != nil {
			fmt.Printf("Task with ID %v not found\n", id)
			return
		}

		newStatus := !task.IsCompleted
		err = database.New(db).UpdateTaskStatus(cmd.Context(), database.UpdateTaskStatusParams{
			ID:          id,
			IsCompleted: newStatus,
			UpdatedAt:   getISTTime(),
		})

		if err != nil {
			log.Fatal("Error updating task: ", err)
		}

		fmt.Printf("Task with ID %s updated to %v\n", id, newStatus)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
