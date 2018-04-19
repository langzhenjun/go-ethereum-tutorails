package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type account struct {
	AddressHex string
	KeyJSON    interface{}
	Password   string
}

// Configs ...
type Configs struct {
	MainAccount account
	Contracts   map[string]string
}

// LoadConfigs ...
func LoadConfigs(path string) *Configs {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load configs file: %v\r\n", err)
	}

	var configs Configs
	err = json.Unmarshal(file, &configs)
	if err != nil {
		log.Fatalf("Failed to decode configs file: %v\r\n", err)
	}

	return &configs
}

//Save ...
func (cfg *Configs) Save() {
	bytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal configs content: %v\r\n", err)
	}

	err = ioutil.WriteFile("../configs.json", bytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write configs file: %v\r\n", err)
	}
}
