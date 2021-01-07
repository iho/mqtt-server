package models

type Message struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
