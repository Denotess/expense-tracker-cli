package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var summaryMonth int

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show total expenses (optionally by month)",
	RunE: func(cmd *cobra.Command, args []string) error {
		total, err := store.Summary(summaryMonth)
		if err != nil {
			return err
		}

		if summaryMonth == 0 {
			fmt.Printf("Total expenses: %s\n", formatMoney(total))
			return nil
		}

		monthName := time.Month(summaryMonth).String()
		fmt.Printf("Total expenses for %s: %s\n", monthName, formatMoney(total))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().IntVar(&summaryMonth, "month", 0, "Month number 1-12 (omit or 0 for all)")
}
