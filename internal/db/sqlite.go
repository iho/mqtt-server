package db

import (
	"database/sql"
	"mqtt-server/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	DB *sql.DB
}

func newSqliteStore(DSN string) (*sqliteStore, error) {
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return nil, err
	}

	return &sqliteStore{
		DB: db,
	}, nil
}

func (store *sqliteStore) AllMessages() ([]models.Message, error) {
	messages := []models.Message{}

	rows, err := store.DB.Query("select * from messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		m := models.Message{}

		err := rows.Scan(&m.ID, &m.Message, &m.Name)
		if err != nil {
			continue
		}

		messages = append(messages, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (store *sqliteStore) GetMessage(id int) (models.Message, error) {
	row := store.DB.QueryRow("select * from messages where id = $1", id)
	message := models.Message{}

	return message, row.Scan(&message.ID, &message.Message, &message.Name)
}

func (store *sqliteStore) InsertMessage(message models.Message) error {
	_, err := store.DB.Exec("insert into messages (message, name) values ($1, $2);", message.Message, message.Name)

	return err
}
