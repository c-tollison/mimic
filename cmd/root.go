package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var connString string

var rootCmd = &cobra.Command{
	Use:          "mimic",
	Short:        "Postgres migration tool",
	SilenceUsage: true,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if connString == "" {
			connString = os.Getenv("MIMIC_DATABASE_URL")
		}

		if connString == "" {
			return fmt.Errorf("no connection string provided: use --conn or set MIMIC_DATABASE_URL")
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&connString, "conn", "c", "", "Postgres connection string (overrides MIMIC_DATABASE_URL)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
