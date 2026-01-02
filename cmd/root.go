package cmd

import (
	"expense-tracker/internal/storage"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	storePath string
	store     *storage.Store
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A simple expense tracker",
	Long:  "expense-tracker lets you add, list, summarize, update, and delete expenses stored locally in a JSON file.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if storePath == "" {
			p, err := storage.DefaultPath()
			if err != nil {
				return err
			}
			storePath = p
		}
		store = storage.New(storePath)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&storePath, "file", "", "Path to expenses JSON file (default is user config dir)")
}
