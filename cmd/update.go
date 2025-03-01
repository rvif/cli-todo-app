package cmd

import (
	"github.com/fatih/color"
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
			color.Red("🔴 Error fetching task: %v", err)
			return
		}

		newStatus := !task.IsCompleted
		err = database.New(db).UpdateTaskStatus(cmd.Context(), database.UpdateTaskStatusParams{
			ID:          id,
			IsCompleted: newStatus,
			UpdatedAt:   getISTTime(),
		})

		strNewStatus := "Pending"
		if newStatus {
			strNewStatus = "Completed"
		}

		if err != nil {
			color.Red("🔴 Error updating task: %v", err)
			return
		}

		color.Green("🟦 Task with ID %s updated to %v", id, strNewStatus)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
