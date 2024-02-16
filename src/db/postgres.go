package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432 // Default port for PostgreSQL
	user     = "yourusername"
	password = "yourpassword"
	dbname   = "yourdbname"
)

type DB interface {
	GetInstance() (DB, error)
}

type postgresDB struct {
	db *sql.DB
}

func (p *postgresDB) GetInstance() (DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	return &postgresDB{
		db: db,
	}, nil
}

func NewPostgresDB() DB {
	return &postgresDB{}
}
