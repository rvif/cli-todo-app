package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		newId := NewShortUUID()

		newTask := Task{
			ID:          newId,
			Name:        taskName,
			IsCompleted: false,
			CreatedAt:   getISTTime(),
			UpdatedAt:   getISTTime(),
		}

		tasks = append(tasks, newTask)
		saveTasks()

		fmt.Printf("Task added: [%s] %s\n", newTask.ID, newTask.Name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
