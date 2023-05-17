package v1

import "github.com/gin-gonic/gin"

func HealthCheck() (int, map[string]interface{}) {
	return 200, gin.H{"status": "UP"}
}
