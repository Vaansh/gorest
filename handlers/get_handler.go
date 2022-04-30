package handlers

import (
	"github.com/gin-gonic/gin"
	data "github.com/vaansh/gorest/data"
	"net/http"
)

func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Items)
}

func GetItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range data.Items {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
