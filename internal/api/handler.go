package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes регистрирует все роуты API
func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/ws", handleWebSocket)
}
