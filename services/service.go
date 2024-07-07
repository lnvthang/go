package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListOfItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Get list of item."})
	}
}

func ReadItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Get item by id."})
	}
}

func CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Create item."})
	}
}

func EditItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Edit Item."})
	}
}

func DeleteItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Delete Item."})
	}
}
