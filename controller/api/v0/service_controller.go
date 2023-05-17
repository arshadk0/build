package v1

import (
	helpers "build/helper/api/v0"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	status, data := helpers.HealthCheck()
	c.JSON(status, data)
}
