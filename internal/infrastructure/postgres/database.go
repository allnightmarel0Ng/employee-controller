package postgres

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type Database struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewDatabase(ctx context.Context, connStr string) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &Database{
		pool: pool,
		ctx:  ctx,
	}, nil
}

func (d *Database) Close() {
	d.pool.Close()
}

func (d *Database) Query(query string, args ...interface{}) (pgx.Rows, error) {
	return d.pool.Query(d.ctx, query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.pool.Exec(d.ctx, query, args...)
}

func (d *Database) QueryRow(query string, args ...interface{}) pgx.Row {
	return d.pool.QueryRow(d.ctx, query, args...)
}

func (d *Database) Begin(ctx context.Context) (pgx.Tx, error) {
	return d.pool.Begin(ctx)
}
