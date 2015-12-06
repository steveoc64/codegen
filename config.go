package main

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

// Runtime variables, held in external file config.json
type ConfigType struct {
	DataSourceName string
}

var Config ConfigType

// Load the config.json file, and override with runtime flags
func _loadConfig() {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(usr.HomeDir)

	// look for codegen.json in the current directory before getting the
	// generic one from the home directory

	cf, err := os.Open("codegen.json")
	if err != nil {
		cf, err = os.Open(usr.HomeDir + "/codegen.json")
		if err != nil {
			log.Println("Could not open ~/codegen.json :", err.Error())
		}
	}

	data := json.NewDecoder(cf)
	if err = data.Decode(&Config); err != nil {
		log.Fatalln("Failed to load ~/codegen.json :", err.Error())
	}
}
