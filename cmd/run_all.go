package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runAllCmd = &cobra.Command{
	Use:   "run-all",
	Short: "Run all pending migrations in order",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented")
	},
}

func init() {
	rootCmd.AddCommand(runAllCmd)
}
