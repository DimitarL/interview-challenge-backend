package storage

import (
	"context"
	"fmt"
	"log"

	pgx "github.com/jackc/pgx/v4"
)

type AppStorage struct {
	conn *pgx.Conn
}

func NewAppStorage() *AppStorage {
	conn, err := pgx.Connect(context.Background(), "postgres://root:root@localhost:5434/testingwithrentals")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	if err == nil {
		fmt.Println("Connected successfully!")
	}

	return &AppStorage{conn: conn}
}

func (a AppStorage) CloseConn() error {
	return a.conn.Close(context.Background())
}
