package service

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Service struct {
	db *bun.DB
}

func New() (*Service, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	return &Service{
		db: bun.NewDB(sqldb, sqlitedialect.New()),
	}, nil
}

func (s *Service) DB() *bun.DB {
	return s.db
}
