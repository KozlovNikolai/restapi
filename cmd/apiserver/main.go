package main

import (
	"flag"
	"log"

	"github.com/KozlovNikolai/restapi/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config) //заполнить конфиг
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)        //создать сервер
	if err := s.Start(); err != nil { //запустить сервер
		log.Fatal(err)
	}
}
