/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task_id]",
	Short: "Delete a todo by its ID",

	Run: func(cmd *cobra.Command, args []string) {

		if len(tasks) == 0 {
			fmt.Println("No tasks to delete.")
			return
		}

		id := args[0]
		found := false
		for i, task := range tasks {
			if string(task.ID) == id {
				found = true
				tasks = append(tasks[:i], tasks[i+1:]...)
				saveTasks()
				break
			}
		}

		if !found {
			fmt.Printf("Task with ID %s not found\n", id)
			return
		} else {
			fmt.Printf("Task with ID %s deleted\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
