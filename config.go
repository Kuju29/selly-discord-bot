package main

import (
	"encoding/json"
	"log"
	"os"
)

func loadConfig() Configuration {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Print("error: Error decoding JSON: " + err.Error())
	}
	return configuration
}
