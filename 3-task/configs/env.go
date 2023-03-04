package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type DataEnv struct {
	ListenPort string `env:"LISTEN_PORT"`
	TimeFormat string `env:"TIME_FORMAT"`
}

func ReadENVConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("ENV File cannot be opened", err)
	}

	data := DataEnv{}
	if err := env.Parse(&data); err != nil {
		log.Fatal("ENV File can not be parsed")
	}
	return &data
}

func (text *DataEnv) PrintConfig() {
	fmt.Printf("Listen Port is: %s\n", text.ListenPort)
	fmt.Printf("Time Format is: %s\n", text.TimeFormat)
}
