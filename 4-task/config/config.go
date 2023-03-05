package config

import (
	"encoding/json"
	"log"
	"os"
)

type DataJson struct {
	ListenPort string `json:ListenPort`
	TimeFormat string `json:TimeFormat`
}

func ReadJsonConfigFile(path string) *DataJson {
	rawData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("JSON file cannot be opened", err)
	}
	var data DataJson
	if err := json.Unmarshal(rawData, &data); err != nil {
		log.Fatal("JSON File cannot be parsed")
	}
	return &data
}
