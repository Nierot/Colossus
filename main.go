package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	var (
		r = gin.Default()
	)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.Set("Verbose", true)

	viper.SetDefault("Port", 8080)
	viper.SetDefault("ImagePath", "/images")
	viper.SetDefault("StaticPath", "/static")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// Create all handlers
	MakeStaticRoutes(r)

	r.Run("0.0.0.0:" + viper.GetString("port"))
}
