package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	jsonData, err := json.MarshalIndent(cfg, "", "  ")

	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	file, err := os.Create(filePath)

	if err != nil {
		return nil
	}
	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}

	file.Sync()

	return nil
}
