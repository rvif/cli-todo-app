package cmd

import (
	"os"

	"github.com/fatih/color"
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
			// Prints to console in red color
			color.Red("🔴 Error fetching tasks: %v", err)
			return
		}
		if len(tasks) == 0 {
			color.Yellow("⭐ No tasks found.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Task Name", "Status", "Created At", "Updated At"})
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgHiMagentaColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.Bold},
		)

		table.SetBorder(false)

		for _, task := range tasks {
			status := color.HiRedString("Pending")
			if task.IsCompleted {
				status = color.HiGreenString("Completed")
			}

			table.Append([]string{
				color.HiMagentaString(task.ID),
				task.Name,
				status,
				task.CreatedAt.Format("02 Jan 06 03:04 PM"),
				task.UpdatedAt.Format("02 Jan 06 03:04 PM"),
			})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
