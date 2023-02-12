package main

import (
	"WB0/internal/config"
	"WB0/internal/nats"
	"WB0/internal/producer"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	natsConn, err := nats.NewNatsConnect(config, logger, "publisher")
	if err != nil {
		logger.Fatalf("failed to connect to nats server: %+v", err)
	}
	ps := producer.NewProducerServer(logger, config, natsConn)

	if err := ps.Run(); err != nil {
		log.Fatalf("failed to start producer server: %+v", err)
	}

}
