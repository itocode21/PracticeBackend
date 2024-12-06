package body

import (
	"net/http"
	"strconv"
	"todoapi/logic"

	"github.com/gin-gonic/gin"
)

func getTodos(contex *gin.Context) {
	page, err := strconv.ParseInt(contex.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, err := strconv.ParseInt(contex.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	todos, err := logic.GetTodos(page, limit)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todoCount := len(todos)
	contex.JSON(http.StatusOK, gin.H{"data": todos, "page": page, "limit": limit, "total": todoCount})
}

func createTodos(contex *gin.Context) {
	userEmail := contex.GetString("email")

	var todo logic.Todo
	err := contex.ShouldBind(&todo)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoCreated, err := todo.CreateTodo(userEmail)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	contex.JSON(http.StatusCreated, todoCreated)
}

func updateTodos(contex *gin.Context) {
	userEmail := contex.GetString("email")

	id, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo logic.Todo
	err = contex.ShouldBind(&todo)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoUpdated, err := todo.UpdateTododbyId(id, userEmail)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contex.JSON(http.StatusOK, todoUpdated)
}

func deleteTodos(contex *gin.Context) {
	userEmail := contex.GetString("email")

	id, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = logic.DeleteTodoById(id, userEmail)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contex.JSON(http.StatusNoContent, nil)
}
