package config

import "database/sql"

type Database interface {
	Open(driverName, dataSourceName string) (*sql.DB, error)
	Close() error
}

type SQLDatabase struct {
	db *sql.DB
}

func NewSQLDatabase() *SQLDatabase {
	return &SQLDatabase{}
}

func (s *SQLDatabase) Open(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	s.db = db
	return db, nil
}

func (s *SQLDatabase) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
