package db

import (
	"mqtt-server/internal/models"
	"sync"
)

type memStore struct {
	messages []models.Message
	lastID   int
	mutex    sync.RWMutex
}

func newMemoryStore() *memStore {
	return &memStore{
		messages: make([]models.Message, 0),
	}
}

func (store *memStore) AllMessages() ([]models.Message, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	return store.messages, nil
}

func (store *memStore) GetMessage(id int) (models.Message, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, message := range store.messages {
		if message.ID == id {
			return message, nil
		}
	}

	return models.Message{}, nil
}

func (store *memStore) InsertMessage(message models.Message) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.lastID++
	message.ID = store.lastID
	store.messages = append(store.messages, message)

	return nil
}
