package logic

import (
	"fmt"
	conndb "todoapi/connDB"

	"todoapi/tools"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	Username string
	Email    string
	Password string
}

// Метод для регистрации пользователя
func (user *User) RegUser() (string, error) {
	db, err := conndb.ConnectDB()
	if db == nil {
		return "", fmt.Errorf("db connection is nil")
	}
	if err != nil {
		fmt.Print("Error user.go 25 string")
	}

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	hashedPassword, err := tools.HashPass(user.Password)
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(user.Username, user.Email, hashedPassword)
	if err != nil {
		return "", err
	}
	return tools.GenerateToken(user.Email)
}

func (user *User) Login() (string, error) {
	db, err1 := conndb.ConnectDB()
	if db == nil {
		return "", fmt.Errorf("db connection is nil")
	}
	if err1 != nil {
		fmt.Print("Error user.go 25 string")
	}
	query := "SELECT * FROM users WHERE email = $2"
	row := db.QueryRow(query, user.Email)

	var dbUser User
	err := row.Scan(&dbUser.Username, &dbUser.Email, &dbUser.Password)
	if err != nil {
		return "", nil
	}

	if !tools.CheckHash(user.Password, dbUser.Password) {
		return "", nil
	}

	return tools.GenerateToken(user.Email)
}
