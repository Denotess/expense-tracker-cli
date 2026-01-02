package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteID int64

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := store.Delete(deleteID); err != nil {
			return err
		}
		fmt.Println("Expense deleted successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int64Var(&deleteID, "id", 0, "Expense ID to delete")

	_ = deleteCmd.MarkFlagRequired("id")
}
