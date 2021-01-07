package server

import (
	"mqtt-server/internal/config"
	"mqtt-server/internal/db"
	"sync"

	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

func TestMessageHandler(t *testing.T) {
	store, _ := db.NewDBStore(db.Memory, "")
	client := &mockClient{}
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	log := logger.Sugar()
	conf := &config.Config{
		MQTTTopic: "test/topic",
	}
	srv := NewServer(store, log, client, conf)
	msg := &message{
		payload: []byte("{\"message\": \"test\", \"name\": \"mike\"}"),
	}
	srv.MessageHandler(client, msg)
	messages, _ := store.AllMessages()
	assert.True(t, len(messages) == 1)
}

func TestMessageHandlerFail(t *testing.T) {
	store, _ := db.NewDBStore(db.Memory, "")
	client := &mockClient{}
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	log := logger.Sugar()
	conf := &config.Config{
		MQTTTopic: "test/topic",
	}
	srv := NewServer(store, log, client, conf)

	msg := &message{
		payload: []byte("{broken json"),
	}
	srv.MessageHandler(client, msg)
	messages, _ := store.AllMessages()
	assert.True(t, len(messages) == 0)

}

type mockClient struct {
}

func (m *mockClient) IsConnected() bool {
	return true
}

func (m *mockClient) IsConnectionOpen() bool {
	return true
}

func (m *mockClient) Connect() mqtt.Token {
	return nil
}

func (m *mockClient) Disconnect(quiesce uint) {
}

func (m *mockClient) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return nil
}

func (m *mockClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return nil
}

func (m *mockClient) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	return nil
}

func (m *mockClient) Unsubscribe(topics ...string) mqtt.Token {
	return nil
}

func (m *mockClient) AddRoute(topic string, callback mqtt.MessageHandler) {}

func (m *mockClient) OptionsReader() mqtt.ClientOptionsReader {
	return mqtt.ClientOptionsReader{}
}

type message struct {
	duplicate bool
	qos       byte
	retained  bool
	topic     string
	messageID uint16
	payload   []byte
	once      sync.Once
	ack       func()
}

func (m *message) Duplicate() bool {
	return m.duplicate
}

func (m *message) Qos() byte {
	return m.qos
}

func (m *message) Retained() bool {
	return m.retained
}

func (m *message) Topic() string {
	return m.topic
}

func (m *message) MessageID() uint16 {
	return m.messageID
}

func (m *message) Payload() []byte {
	return m.payload
}

func (m *message) Ack() {
	m.once.Do(m.ack)
}
