package postgres

import (
	"context"
	"fmt"

	"github.com/Team-Hype/vault-secret-management/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
)

type DBInfo struct {
	pool *pgxpool.Pool
}

func NewDBInfo(pool *pgxpool.Pool) *DBInfo {
	return &DBInfo{pool: pool}
}

const DBInfoScan = `
SELECT
	key::TEXT,
	value::TEXT
FROM db_info
`

func (d *DBInfo) Scan(ctx context.Context) ([]*model.Info, error) {
	rows, err := d.pool.Query(ctx, DBInfoScan)
	if err != nil {
		return nil, fmt.Errorf("db_info repository: query error: %w", err)
	}
	defer rows.Close()

	infos, err := pgx.CollectRows[*model.Info](rows, pgx.RowToAddrOfStructByName)
	if err != nil {
		return nil, fmt.Errorf("db_info repository: collect rows error: %w", err)
	}

	return infos, nil
}
