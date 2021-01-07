package db

import (
	"fmt"
	"mqtt-server/internal/models"
)

type Type int

const (
	Memory Type = iota
	Sqlite
)

type DBStore interface {
	AllMessages() ([]models.Message, error)
	GetMessage(id int) (models.Message, error)
	InsertMessage(message models.Message) error
}

func NewDBStore(storeType Type, DSN string) (DBStore, error) {
	switch storeType {
	case Sqlite:
		return newSqliteStore(DSN)
	case Memory:
		return newMemoryStore(), nil
	}

	return nil, fmt.Errorf("kind has to be either sqlite or memory")
}
