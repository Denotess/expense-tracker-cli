package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updateID          int64
	updateAmount      int64
	updateDescription string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an expense by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := store.Update(updateID, updateAmount, updateDescription); err != nil {
			return err
		}
		fmt.Println("Expense updated successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().Int64Var(&updateID, "id", 0, "Expense ID to update")
	updateCmd.Flags().Int64Var(&updateAmount, "amount", 0, "New amount in cents (e.g. 2000 = $20.00)")
	updateCmd.Flags().StringVar(&updateDescription, "description", "", "New description")

	_ = updateCmd.MarkFlagRequired("id")
	_ = updateCmd.MarkFlagRequired("amount")
	_ = updateCmd.MarkFlagRequired("description")
}
