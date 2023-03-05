package main

import (
	"fmt"
	"testing_task/configs"
)

func main() {
	fmt.Println("JSON Config")
	configJSON := configs.ReadJsonConfigFile("configs/data/data.json")
	configJSON.PrintConfig()

	fmt.Println("YAML Config")
	configYAML := configs.ReadYAMLConfigFile("configs/data/data.yaml")
	configYAML.PrintConfig()

	fmt.Println("Flags Config")
	configFlags := configs.ReadFlags()
	configFlags.PrintConfig()

	fmt.Println("ENV Config")
	configENV := configs.ReadENVConfig("configs/data/data.env")
	configENV.PrintConfig()
}
