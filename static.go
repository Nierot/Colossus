package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func MakeStaticRoutes(r *gin.Engine) {
	group := r.Group(viper.GetString("StaticPath"))

	makeImageRoute(group)
}

func makeImageRoute(r *gin.RouterGroup) {
	r.GET(viper.GetString("ImagePath")+"/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
}
