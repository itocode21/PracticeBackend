package conndb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// подключение к базе данных PostgreSQL
func ConnectDB() (*sql.DB, error) {

	connStr := "postgres://ito21:1899@localhost:5432/TODO"

	ctx := context.Background()

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии подключения: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err := createTable(db); err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ошибка при проверке подключения: %w", err)
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	createTodosTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		description TEXT,
		createdBy TEXT
		);`

	_, err := db.Exec(createTodosTable)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы todos: %w", err)
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT,
		email TEXT,
		password TEXT
	);
	`
	_, err = db.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы users: %w", err)
	}

	return nil
}
