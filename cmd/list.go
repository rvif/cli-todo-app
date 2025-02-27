package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/rvif/cli-todo-app/internal/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list todos",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := database.New(db).GetAllTasks(cmd.Context())
		if err != nil {
			log.Fatalf("Error fetching tasks: %v", err)
		}
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

			table.Append([]string{
				task.ID,
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
