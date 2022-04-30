package handlers

import (
	"github.com/gin-gonic/gin"
	data "github.com/vaansh/gorest/data"
	"net/http"
)

func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Items)
}
