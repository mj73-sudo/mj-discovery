package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	LoadEnv()
	LoadLogConfig()
}

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: AllowOrigins,
		AllowMethods: AllowMethods,
		AllowHeaders: AllowHeaders,
	}))
	r.Run(":" + ApplicationPort)
}
