package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dbychkar/go_chat/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Отключаем CORS для простоты
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		log.Printf("Received: %+v", msg)

		// Пока просто возвращаем обратно (echo)
		response := fmt.Sprintf("Echo from server: %s", msg.Content)
		if err := conn.WriteJSON(models.Message{
			Sender:  "server",
			Content: response,
		}); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
