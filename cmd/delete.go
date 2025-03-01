package cmd

import (
	"context"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task_id]",
	Short: "Delete a todo by its ID",

	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		res, err := db.ExecContext(context.Background(), `DELETE FROM tasks WHERE id = $1`, id)
		if err != nil {
			color.Red("🔴 Error deleting task: %v", err)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			color.Red("🟡 Task with ID %s not found", id)
			return
		}

		color.Green("🟦 Task with ID %s deleted", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
