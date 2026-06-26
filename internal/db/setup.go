package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (c *Client) Setup(ctx context.Context, schema string) error {
	schemaExists, err := c.SchemaExists(ctx, schema)
	if err != nil {
		return err
	}
	if !schemaExists {
		return fmt.Errorf("schema %q does not exist", schema)
	}

	tablesExist, err := c.TablesExist(ctx, schema)
	if err != nil {
		return err
	}
	if tablesExist {
		return fmt.Errorf("mimic is already set up in schema %q — run 'mimic status' to see migration state", schema)
	}

	return c.RunInTx(ctx, func(tx pgx.Tx) error {
		statements := []string{
			fmt.Sprintf(`CREATE TYPE %q.mimic_state AS ENUM ('running', 'applied', 'failed', 'rolled_back')`, schema),
			fmt.Sprintf(`CREATE TABLE %q.mimic_migrations (
				id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				version     TEXT NOT NULL,
				name        TEXT NOT NULL,
				state       %q.mimic_state NOT NULL,
				occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				error       TEXT
			)`, schema, schema),
			fmt.Sprintf(`CREATE INDEX ON %q.mimic_migrations (version, occurred_at DESC)`, schema),
			fmt.Sprintf(`CREATE TABLE %q.mimic_scripts (
				id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				name        TEXT NOT NULL,
				state       %q.mimic_state NOT NULL,
				occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				error       TEXT
			)`, schema, schema),
			fmt.Sprintf(`CREATE INDEX ON %q.mimic_scripts (name, occurred_at DESC)`, schema),
		}

		for _, sql := range statements {
			if _, err := tx.Exec(ctx, sql); err != nil {
				return err
			}
		}
		return nil
	})
}
