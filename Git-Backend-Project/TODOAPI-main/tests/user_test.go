package tests

import (
	"fmt"
	"testing"
	conndb "todoapi/connDB"
	"todoapi/logic"
)

// find nill pointer ned fix
func TestRegUser(t *testing.T) {
	db, err := conndb.ConnectDB()
	if db == nil {
		t.Fatal("Database connection is nil")
	}
	if err != nil {
		fmt.Print("Error TestRegUser")
	}

	user := logic.User{Username: "TestName", Email: "TestEmail", Password: "TestPassword"}
	token, err := user.RegUser()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected a token, got empty string")
	}
}
