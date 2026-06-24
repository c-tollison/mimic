package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mimic",
	Short: "Postgres migration tool with multi-tenant schema support",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Mimic '%s' \n", err)
		os.Exit(1)
	}
}
