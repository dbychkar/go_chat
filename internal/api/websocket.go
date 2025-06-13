package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dbychkar/go_chat/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var hub = NewHub()

func init() {
	go hub.Run()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	hub.Register <- client

	// Отправка истории сообщений при подключении
	for _, msg := range hub.Storage.GetAll() {
		data, err := json.Marshal(msg)
		if err == nil {
			client.Send <- data
		}
	}

	go writePump(client)
	readPump(client)
}

func readPump(client *Client) {
	defer func() {
		hub.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, data, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}

		var incoming models.Message
		if err := json.Unmarshal(data, &incoming); err != nil {
			continue
		}

		// ⬇️ Создаём сообщение с добавленным временем
		msg := models.Message{
			Username:  incoming.Username,
			Text:      incoming.Text,
			Timestamp: time.Now().Format("15:04:05"),
		}

		// Сохраняем и рассылаем сообщение
		hub.Storage.Add(msg)

		encoded, err := json.Marshal(msg)
		if err == nil {
			hub.Broadcast <- encoded
		}
	}
}

func writePump(client *Client) {
	defer client.Conn.Close()

	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
