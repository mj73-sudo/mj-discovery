package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetGin() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: AllowOrigins,
		AllowMethods: AllowMethods,
		AllowHeaders: AllowHeaders,
	}))
	LogJson.WithFields(logrus.Fields{"AllowOrigin": AllowOrigins, "AlloMethods": AllowMethods, "AllowHeaders": AllowHeaders}).Debugln("Cors configuration")
	return r
}
