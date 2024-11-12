package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) {
	c.CurrentUserName = user
	err := write(*c)
	if err != nil {
		log.Fatal(err)
	}
}

func Read() Config {
	filePath, err := getConfigFilePath()

	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	config := Config{}

	err = json.Unmarshal(jsonData, &config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}
