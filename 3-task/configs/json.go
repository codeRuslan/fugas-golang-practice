package configs

import (
	"encoding/json"
	"fmt"
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
		log.Fatal("JSON File cannot be opened", err)
	}
	var data DataJson
	if err := json.Unmarshal(rawData, &data); err != nil {
		log.Fatal("JSON File can not be parsed")
	}
	return &data
}

func (text *DataJson) PrintConfig() {
	fmt.Printf("Listen Port is: %s\n", text.ListenPort)
	fmt.Printf("Time Format is: %s\n", text.TimeFormat)
}
