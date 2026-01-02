package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addDescription string
	addAmount      int64
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	RunE: func(cmd *cobra.Command, args []string) error {
		exp, err := store.Add(addAmount, addDescription)
		if err != nil {
			return err
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", exp.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&addDescription, "description", "", "Expense description")
	addCmd.Flags().Int64Var(&addAmount, "amount", 0, "Expense amount")

	_ = addCmd.MarkFlagRequired("description")
	_ = addCmd.MarkFlagRequired("amount")
}
