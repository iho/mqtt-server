package main

import (
	"fmt"
	"log"
	"mqtt-server/internal/config"
	"mqtt-server/internal/db"
	"mqtt-server/internal/server"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

func main() {
	conf, err := config.ReadConfig("./conf.json")
	if err != nil {
		log.Fatal(err)
	}

	store, err := db.NewDBStore(db.Sqlite, conf.DSN)
	if err != nil {
		log.Fatal(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync() // nolint: errcheck

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", conf.Broker, conf.Port))

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Fatal(token.Error().Error())
	}

	srv := server.NewServer(store, logger.Sugar(), client, conf)
	srv.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
