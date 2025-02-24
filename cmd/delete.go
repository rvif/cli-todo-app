package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task_id]",
	Short: "Delete a task by ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := ShortUUID(args[0])

		index := -1
		for i, task := range tasks {
			if task.ID == taskID {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Printf("Task with ID %s not found.\n", taskID)
			return
		}

		tasks = append(tasks[:index], tasks[index+1:]...)
		saveTasks()

		fmt.Printf("Task [%s] deleted successfully!\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
