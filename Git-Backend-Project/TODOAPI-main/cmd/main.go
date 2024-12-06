package main

import (
	"todoapi/body"
	conndb "todoapi/connDB"

	"github.com/gin-gonic/gin"
)

func main() {
	conndb.ConnectDB()
	server := gin.Default()
	body.RegisterServer(*server)
	server.Run(":8080")

}
