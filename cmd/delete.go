package cmd

import (
	"context"
	"fmt"
	"log"

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
			log.Fatalf("Error deleting task: %v", err)
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			fmt.Printf("Task with ID %s not found\n", id)
			return
		}

		fmt.Printf("Task with ID %s deleted\n", id)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
