package main

import (
	"WB0/internal/apiserver"
	"WB0/internal/config"
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
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
