package mid

import (
	"net/http"
	"todoapi/tools"

	"github.com/gin-gonic/gin"
)

func Authentication(contex *gin.Context) {
	token := contex.GetHeader("Authorization")
	if token == "" {
		contex.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	email, err := tools.ValidateToken(token)
	if err != nil {
		contex.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	contex.Set("email", email)
	contex.Next()
}
