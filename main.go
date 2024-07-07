package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"module-go/services"
	// "module-go/database"
)

func main() {
	r := gin.Default()
	// db := database.StartPostgres()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api")
	{
		v1.GET("/items", services.GetListOfItems())
		v1.GET("/item/:id", services.ReadItemById())
		v1.POST("/item", services.CreateItem())
		v1.PUT("/item/:id", services.EditItemById())
		v1.DELETE("/item/:id", services.DeleteItemById())
	}

	r.Run(":8080")
}
