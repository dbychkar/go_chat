package storage

import (
	"sync"

	"github.com/dbychkar/go_chat/models"
)

type MemoryStorage struct {
	messages []models.Message
	mu       sync.Mutex
	limit    int
}

func NewMemoryStorage(limit int) *MemoryStorage {
	return &MemoryStorage{
		messages: make([]models.Message, 0, limit),
		limit:    limit,
	}
}

func (s *MemoryStorage) Add(msg models.Message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.messages) >= s.limit {
		s.messages = s.messages[1:]
	}
	s.messages = append(s.messages, msg)
}

func (s *MemoryStorage) GetAll() []models.Message {
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]models.Message(nil), s.messages...)
}
