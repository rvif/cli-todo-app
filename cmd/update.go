/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [task_id]",
	Short: "Update a todo by its ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tasks) == 0 {
			fmt.Println("No tasks to update.")
			return
		}

		id := args[0]
		found := false
		updatedTask := Task{}
		for i, _ := range tasks {
			if string(tasks[i].ID) == id {
				found = true

				tasks[i].IsCompleted = !tasks[i].IsCompleted
				tasks[i].UpdatedAt = getISTTime()

				updatedTask = tasks[i]
				saveTasks()
				break
			}
		}

		var status string
		if updatedTask.IsCompleted {
			status = "completed"
		} else {
			status = "remaining"
		}

		if !found {
			fmt.Printf("Task with ID %s not found\n", id)
			return
		} else {
			fmt.Printf("Task with ID %s updated to %v\n", id, status)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
