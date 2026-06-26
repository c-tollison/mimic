package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Client struct {
	conn *pgx.Conn
}

func New(ctx context.Context, connString string) (*Client, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		conn.Close(ctx)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.conn.Close(ctx)
}

func (c *Client) RunInTx(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := c.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx failed: %w; rollback also failed: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
