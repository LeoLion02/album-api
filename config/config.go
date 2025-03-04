package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type ConnectionString struct {
	SqlServer string `json:"sqlServer"`
}

type AppConfig struct {
	ConnectionString ConnectionString `json:"connectionStrings"`
}

var (
	configInstance *AppConfig
	once           sync.Once
)

func LoadConfig() (*AppConfig, error) {
	once.Do(func() {
		configFilePath := os.Getenv("CONFIG_FILE")

		if configFilePath == "" {
			_, filename, _, _ := runtime.Caller(0)
			configFilePath = filepath.Dir(filename) + `/config.json`
		}

		fileData, err := os.ReadFile(configFilePath)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		err = json.Unmarshal(fileData, &configInstance)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			configInstance = nil
		}
	})

	if configInstance == nil {
		return nil, fmt.Errorf("failed to load configuration")
	}

	return configInstance, nil
}

func GetConfig() (*AppConfig, error) {
	return LoadConfig()
}
