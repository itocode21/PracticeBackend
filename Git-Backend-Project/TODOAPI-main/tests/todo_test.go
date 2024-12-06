package tests

import (
	"log"
	"testing"
	conndb "todoapi/connDB"
	"todoapi/logic"
)

func TestCreateTodo(t *testing.T) {
	conndb.ConnectDB()
	todo := logic.Todo{Title: "Test todo", Description: "Test Desc"}
	createdTodo, err := todo.CreateTodo("test@exmaple.com")
	if err != nil {
		log.Fatalf("Expected no error, got %v", err)
	}

	if createdTodo.Title != todo.Title {
		t.Errorf("Expexted title %s,got %s", todo.Title, createdTodo.Title)
	}

}

func TestGetTodos(t *testing.T) {
	conndb.ConnectDB()
	todos, err := logic.GetTodos(1, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(todos) == 0 {
		t.Errorf("Expected at least one todo")
	}
}
