package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func ConnectBD() {
	// Строка подключения к вашей базе данных
	connString := fmt.Sprintf(
		"host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Подключение к базе данных
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Successfully connected to the database!")
}
