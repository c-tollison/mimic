package cmd

import (
	"context"
	"fmt"

	"github.com/c-tollison/mimic/internal/db"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup <schema>",
	Short: "Initialize mimic tracking tables in the target schema",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		schema := args[0]
		ctx := context.Background()

		client, err := db.New(ctx, connString)
		if err != nil {
			return fmt.Errorf("could not connect to database: %w", err)
		}
		defer client.Close(ctx)

		fmt.Printf("Setting up mimic in schema %q...\n", schema)

		if err := client.Setup(ctx, schema); err != nil {
			return err
		}

		fmt.Printf("Setup complete.\n")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
