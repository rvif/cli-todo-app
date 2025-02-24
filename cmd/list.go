package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tasks) == 0 {
			fmt.Println("No tasks found. noice.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Task Name", "Status", "Created At", "Updated At"})

		for _, task := range tasks {
			status := "remaining"
			if task.IsCompleted {
				status = "completed"
			}

			StrShortUUID := string(task.ID)

			table.Append([]string{
				StrShortUUID,
				task.Name,
				status,
				task.CreatedAt.Format(time.RFC822),
				task.UpdatedAt.Format(time.RFC822),
			})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
