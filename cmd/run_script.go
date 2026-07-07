package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runScriptCmd = &cobra.Command{
	Use:   "run-script <name>",
	Short: "Run a tracked one-off script",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented")
	},
}

func init() {
	rootCmd.AddCommand(runScriptCmd)
}
