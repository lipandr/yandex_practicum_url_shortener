package dao

import (
	"database/sql"

	_ "github.com/jackc/pgx/stdlib"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	if _, err = db.Exec(URLTable); err != nil {
		return nil, err
	}

	return db, nil
}
