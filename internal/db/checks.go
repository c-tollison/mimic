package db

import (
	"context"
	"fmt"
)

func (c *Client) SchemaExists(ctx context.Context, schema string) (bool, error) {
	var exists bool
	err := c.conn.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM information_schema.schemata
			WHERE schema_name = $1
		)
	`, schema).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check schema existence: %w", err)
	}
	return exists, nil
}

func (c *Client) TablesExist(ctx context.Context, schema string) (bool, error) {
	var count int
	err := c.conn.QueryRow(ctx, `
		SELECT COUNT(*) FROM information_schema.tables
		WHERE table_schema = $1
		AND table_name IN ('mimic_migrations', 'mimic_scripts')
	`, schema).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check existing tables: %w", err)
	}
	return count > 0, nil
}
