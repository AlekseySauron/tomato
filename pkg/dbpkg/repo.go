package dbpkg

import "database/sql"

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) DBRepository {
	return &SQLRepository{
		db: db,
	}
}

func (r *SQLRepository) Close() error {
	return r.db.Close()
}
