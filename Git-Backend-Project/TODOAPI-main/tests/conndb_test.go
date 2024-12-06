package tests

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	conndb "todoapi/connDB"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// TestConnDB проверяет, что функция ConnDB успешно открывает соединение с базой данных и создает таблицу users.
func TestConnDB(t *testing.T) {
	// тестовая строка подклчения к бд
	connStr := "postgres://ito21:1899@localhost:5432/TODO"

	// Открываем соединение с базой данных
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Вызываем функцию ConnDB
	createdDB, err := conndb.ConnectDB()

	// Проверяем, что соединение, возвращаемое ConnDB, не nil
	if createdDB == nil {
		t.Fatal("Expected non-nil database connection")
	}

	if err != nil {
		fmt.Print("Error call connectDB", err)
	}

	// проверка действительно ли соеденение установленно
	err = createdDB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	// Проверяем, что таблица users была создана
	var exists bool
	users := createdDB.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users')").Scan(&exists)
	if users != nil {
		t.Fatalf("Failed to check if table exists: %v", users)
	}

	if !exists {
		t.Fatal("Table users does not exist")
	}

	todos := createdDB.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'todos')").Scan(&exists)
	if todos != nil {
		t.Fatalf("Failed to check if table exists: %v", todos)
	}

	if !exists {
		t.Fatal("Table todos  does not exist")
	}

	log.Println("Test passed: Database connection established and table users, todos created successfully!")
}
