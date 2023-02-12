package main

import (
	"WB0/internal/apiserver"
	"WB0/internal/config"
	"WB0/internal/nats"
	"flag"
	"github.com/BurntSushi/toml"
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
		log.Fatalf("failed to read config file: %+v", err)
	}

	logger, err := apiserver.NewLogger(config)
	if err != nil {
		log.Fatalf("failed to create logger: %+v", err)
	}

	natsConn, err := nats.NewNatsConnect(config, logger, "subscriber")
	if err != nil {
		log.Fatalf("failed to connect to nats server: %+v", err)
	}

	db, err := apiserver.NewDb(config)
	if err != nil {
		log.Fatalf("failed to connect to DB: %+v", err)
	}
	defer db.Close()

	server := apiserver.NewApiServer(
		logger,
		config,
		natsConn,
		db,
	)
	if err := server.Start(); err != nil {
		log.Fatalf("failed to start server: %+v", err)
	}
}
