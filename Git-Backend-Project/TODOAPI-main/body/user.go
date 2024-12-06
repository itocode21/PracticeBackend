package body

import (
	"net/http"
	"todoapi/logic"

	"github.com/gin-gonic/gin"
)

func createUser(contex *gin.Context) {
	var user logic.User
	err := contex.ShouldBind(&user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.RegUser()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	contex.JSON(http.StatusCreated, gin.H{"token": token})
}

func login(contex *gin.Context) {
	var user logic.User
	err := contex.ShouldBind(&user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.Login()
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contex.JSON(http.StatusOK, gin.H{"token": token})
}
