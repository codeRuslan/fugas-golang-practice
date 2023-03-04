package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DataYaml struct {
	ListenPort string `yaml:"ListenPort"`
	TimeFormat string `yaml:"TimeFormat"`
}

func ReadYAMLConfigFile(path string) *DataYaml {
	rawData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("YAML File cannot be opened", err)
	}
	var data DataYaml
	if err := yaml.Unmarshal(rawData, &data); err != nil {
		log.Fatal("YAML File can not be parsed")
	}
	return &data
}

func (text *DataYaml) PrintConfig() {
	fmt.Printf("Listen Port is: %s\n", text.ListenPort)
	fmt.Printf("Time Format is: %s\n", text.TimeFormat)
}
