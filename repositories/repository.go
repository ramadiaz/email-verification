package repositories

import "database/sql"

type CompRepositories interface {
}

type compRepositories struct {
	DB *sql.DB
}

func NewCompRepositories(DB *sql.DB) *compRepositories {
	return &compRepositories{
		DB: DB,
	}
}