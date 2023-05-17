package router

import (
	apiContollersV1 "build/controller/api/v0"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func addHeaders(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
	} else {
		c.Next()
	}
}

func init() {
	Router = gin.Default()
	Router.Use(addHeaders)

	v1 := Router.Group("api/v1")
	{
		v1.GET("/health-check", apiContollersV1.HealthCheck)
	}
}