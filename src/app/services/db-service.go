package services

import (
	"database/sql"
	"fmt"
	config "webserver/src/config"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseService struct {
	db *sql.DB
}

func NewDatabaseService(cnf config.Config) (*DatabaseService, error) {
	// Create a DSN (Data Source Name) string
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cnf.Username, cnf.Password, cnf.Host, cnf.DBName)

	// Connect to the MariaDB database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("connected")

	return &DatabaseService{db}, nil
}

func (ds *DatabaseService) Close() error {
	return ds.db.Close()
}
