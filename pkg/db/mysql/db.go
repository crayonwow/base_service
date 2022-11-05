package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dns = "%s:%s@tcp(%s:%s)/%s&parseTime=true"

func newDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(dns, cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}
	return db, err
}
