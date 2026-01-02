package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/spf13/cobra"
)

func formatMoney(dollars int64) string {
	return fmt.Sprintf("$%d", dollars)
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	RunE: func(cmd *cobra.Command, args []string) error {
		expenses, err := store.List()
		if err != nil {
			return err
		}

		sort.Slice(expenses, func(i, j int) bool {
			if expenses[i].Date.Equal(expenses[j].Date) {
				return expenses[i].ID < expenses[j].ID
			}
			return expenses[i].Date.Before(expenses[j].Date)
		})

		if len(expenses) == 0 {
			fmt.Println("No expenses found.")
			return nil
		}

		fmt.Println("ID  Date        Description                     Amount")
		fmt.Println("-------------------------------------------------------")

		for _, e := range expenses {
			dateStr := e.Date.In(time.Local).Format("2006-01-02")
			fmt.Printf("%-3d %-10s %-30s %s\n",
				e.ID,
				dateStr,
				truncate(e.Description, 30),
				formatMoney(e.Amount),
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
