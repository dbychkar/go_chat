package main

import (
	"github.com/dbychkar/go_chat/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 👇 Статические файлы (отдаём всё из ./web)
	router.Static("/static", "./web")

	// 👇 WebSocket-эндпоинт
	router.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})
	router.GET("/ws", api.HandleWebSocket)

	router.Run(":8080")
}
