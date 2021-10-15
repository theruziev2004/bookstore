package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type Database struct {
	DSN     string
	Konnect *pgx.Conn
}

func NewDatabase(dsn string) *Database {
	return &Database{
		DSN: dsn,
	}
}

func (d *Database) Connect() {
	fmt.Println("Start server...")
	connect, err := pgx.Connect(context.Background(), d.DSN)
	if err != nil {
		log.Fatalf("can't connect to db %e", err)
	}
	d.Konnect = connect
	d.initTables()

}

func (d *Database) initTables() {
	tables := []string{
		userTable,
		bookTable,
		TransactionTable,
	}

	for _, t := range tables {
		_, err := d.Konnect.Exec(context.Background(), t)
		if err != nil {
			log.Fatal("can't create a table", err)
		}
	}
}
