package logic

import (
	"errors"
	"log"
	conndb "todoapi/connDB"
)

type Todo struct {
	Id          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Createdby   string
}

func GetTodos(page, limit int64) ([]Todo, error) {
	db, err := conndb.ConnectDB()
	if db == nil {
		log.Fatal("db is nil!")
	}
	if err != nil {
		log.Fatal("error in db")
	}
	query := `SELECT * FROM todos LIMIT $1 OFFSET $2`
	rows, err := db.Query(query, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Createdby)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (todo *Todo) CreateTodo(userEmail string) (*Todo, error) {
	db, err := conndb.ConnectDB()
	if db == nil {
		log.Fatal("db is nil")
	}
	if err != nil {
		log.Fatal("err in db")
	}

	query := `INSERT INTO todos (title, description, createdby) VALUES ($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Title, todo.Description, userEmail)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	todo.Id = id
	todo.Createdby = userEmail

	return todo, nil
}

func (todo *Todo) UpdateTododbyId(id int64, userEmail string) (*Todo, error) {

	db, err := conndb.ConnectDB()
	if db == nil {
		log.Fatal("db is nil")
	}
	if err != nil {
		log.Fatal("error in db")
	}
	var createdby string
	queryCheck := `SELECT createdby FROM todos WHERE id = $1`
	error := db.QueryRow(queryCheck, id).Scan(&createdby)
	if error != nil {
		return nil, err
	}

	if createdby != userEmail {
		return nil, errors.New("unauth to update this")
	}

	query := `UPDATE todos SET title = $1, description = $2 WHERE id = $3`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description, id)
	if err != nil {
		return nil, err
	}

	todo.Id = id
	todo.Createdby = userEmail
	return todo, nil
}

func DeleteTodoById(id int64, userEmail string) error {
	db, error := conndb.ConnectDB()
	if db == nil {
		log.Fatal("db is nil")
	}
	if error != nil {
		log.Fatal("some error in db")
	}

	var createdby string
	queryCheck := `SELECT createdby FROM todos WHERE id = $1`
	err := db.QueryRow(queryCheck, id).Scan(&createdby)
	if err != nil {
		return err
	}

	if createdby != userEmail {
		return errors.New("unauth to delete this")
	}

	query := `DELETE FROM todos WHERE id = $1`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
