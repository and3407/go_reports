package server

import (
	"github.com/and3407/go_reports/app/http/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	ginEngine := gin.Default()

	routes.InitRoutes(ginEngine)

	ginEngine.Run(port)
}