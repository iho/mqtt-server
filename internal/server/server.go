package server

import (
	"encoding/json"
	"mqtt-server/internal/config"
	"mqtt-server/internal/db"
	"mqtt-server/internal/models"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type Server struct {
	store  db.DBStore
	logger *zap.SugaredLogger
	client mqtt.Client
	config *config.Config
}

func NewServer(store db.DBStore, logger *zap.SugaredLogger, client mqtt.Client, config *config.Config) *Server {
	return &Server{
		store:  store,
		logger: logger,
		client: client,
		config: config,
	}
}

func (s *Server) Run() {
	token := s.client.Subscribe(s.config.MQTTTopic, 1, s.messageHandler)
	token.Wait()
}

func (s *Server) messageHandler(client mqtt.Client, msg mqtt.Message) {
	s.logger.Debugf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	message := models.Message{}

	if err := json.Unmarshal(msg.Payload(), &message); err != nil {
		s.logger.Error(err)
	}

	if err := s.store.InsertMessage(message); err != nil {
		s.logger.Error(err)
	}
	s.logger.Debug("Message succesfully inserted")
}
