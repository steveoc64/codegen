package main

import (
	"encoding/json"
	"log"
	"os"
)

// Runtime variables, held in external file config.json
type ConfigType struct {
	Debug          bool
	DataSourceName string
}

var Config ConfigType

// Load the config.json file, and override with runtime flags
func _loadConfig() {
	cf, err := os.Open("~/codegen.json")
	if err != nil {
		log.Println("Could not open ~/codegen.json :", err.Error())
	}

	data := json.NewDecoder(cf)
	if err = data.Decode(&Config); err != nil {
		log.Fatalln("Failed to load ~/codegen.json :", err.Error())
	}
}
