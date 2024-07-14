package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "module-go/config"
	"module-go/services"
)

func main() {
	r := gin.Default()

	database.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api")
	{
		v1.GET("/users", services.GetUserList)
		v1.GET("/user/:id", services.GetUserById)
		v1.POST("/user", services.CreateUser)
		v1.PUT("/user/:id", services.EditUserById)
		v1.DELETE("/item/:id", services.DeleteItemById())
	}

	r.Run(":8080")
}
