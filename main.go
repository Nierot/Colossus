package main

import (
	"fmt"

	"github.com/Nierot/crm/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	var r = gin.Default()

	makeConfig()

	// Add CORS and Logger middleware
	r.Use(gin.Logger())
	r.Use(cors.Default())
	gin.SetMode(viper.GetString("GinMode"))

	// Create all handlers
	MakeStaticRoutes(r)

	r.GET("/tally/:id", routes.GetTallyList)
	r.DELETE("/tally/:id", routes.DeleteTallyList)
	r.POST("/tally/:id", routes.Tally)

	// Now run the server
	r.Run("0.0.0.0:" + viper.GetString("port"))
}

func makeConfig() {
	// Read in config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.Set("Verbose", true)

	viper.SetDefault("Port", 8080)
	viper.SetDefault("ImagePath", "/images")
	viper.SetDefault("StaticPath", "/static")
	viper.SetDefault("GinMode", "debug")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
