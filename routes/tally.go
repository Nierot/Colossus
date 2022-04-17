package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTallyList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetTallyList",
		"id":      c.Params.ByName("id"),
	})
}

func DeleteTallyList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteTallyList",
		"id":      c.Params.ByName("id"),
	})
}

func Tally(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Tally",
		"id":      c.Params.ByName("id"),
	})
}
