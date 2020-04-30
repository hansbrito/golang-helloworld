package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config é o tipo que recebe o arquivo de configuração
type Config struct {
	Port       int      `json:"port"`
	Technology string   `json:"technology"`
	Redirect   []string `json:"redirect"`
}

//LoadConfiguration Importa a configuração do arquivo
func LoadConfiguration() Config {
	var config Config
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
