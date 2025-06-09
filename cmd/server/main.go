package main

import (
	"log"

	"github.com/dbychkar/go_chat/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Подключаем хендлеры
	api.RegisterRoutes(r)

	log.Println("🚀 Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to run server: %v", err)
	}
}
