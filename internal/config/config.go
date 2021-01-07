package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DSN       string `json:"DSN"`
	MQTTTopic string `json:"topic"`
	Port      int    `json:"port"`
	Broker    string `json:"broker"`
}

func ReadConfig() (*Config, error) {
	file, err := os.Open("./conf.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}

	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
