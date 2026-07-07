package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fillTestDataCmd = &cobra.Command{
	Use:   "fill-test-data",
	Short: "Run a test data script against a target",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented")
	},
}

func init() {
	rootCmd.AddCommand(fillTestDataCmd)
}
