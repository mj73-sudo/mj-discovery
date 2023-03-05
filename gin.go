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

	handleRoute(r)

	return r
}

func handleRoute(r *gin.Engine) {
	r.POST("/api/mjdis/register", addNewService)
	r.POST("/api/mjdis/remove", removeService)
	r.GET("/api/mjdis/list", removeService)
}

func addNewService(c *gin.Context) {
	var newService Service
	if err := c.ShouldBindJSON(&newService); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	newService.SaveNewService()
	c.JSON(200, gin.H{"message": "OK"})
}

func removeService(c *gin.Context) {
	var newService Service
	if err := c.ShouldBindJSON(&newService); err != nil || newService.Name == "" {
		c.JSON(400, gin.H{"message": err.Error()})
	}
	newService.SaveNewService()
	c.JSON(200, gin.H{"message": "OK"})
}

func listServices(c *gin.Context) {

}
