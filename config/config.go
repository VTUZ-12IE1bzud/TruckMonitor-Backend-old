package config

import (
	"os"
	"encoding/json"
	"fmt"
)

/** Конфигурация приложения. */
type AppConfiguration struct {
	DataBaseUrl      string    `json:"dbUrl"`
	DataBaseName     string    `json:"dbName"`
	DataBaseUser     string    `json:"dbUserName"`
	DataBasePassword string    `json:"dbUserPassword"`
}

var appConfig *AppConfiguration

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&appConfig)
	if err != nil {
		fmt.Println("Оишбка чтения конфигурационного файла", err)
		panic("Config Error")
	}
}

func GetAppConfiguration() AppConfiguration {
	return *appConfig
}
