package healthcheck

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	Repo struct {
		db *sqlx.DB
	}
)

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r Repo) Ping() error {
	err := r.db.Ping()
	if err != nil {
		return fmt.Errorf("failed ping psql")

	}
	return nil
}
