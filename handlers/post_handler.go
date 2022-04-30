package handlers

import (
	"github.com/gin-gonic/gin"
	data "github.com/vaansh/gorest/data"
	"net/http"
)

func PostItem(c *gin.Context) {
	var newItem data.Item

	if err := c.BindJSON(&newItem); err != nil {
		panic(err)
	}

	data.Items = append(data.Items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}
