package config

import (
	"os"
	"encoding/json"
	"fmt"
)

/** Конфигурация приложения. */
type AppConfiguration struct {
	DBUrl      string    `json:"dbUrl"`
	DBName     string    `json:"dbName"`
	DBUser     string    `json:"dbUserName"`
	DBPassword string    `json:"dbUserPassword"`
	TokenKey   string    `json:"tokenKey"`
}

var conf *AppConfiguration

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Оишбка чтения конфигурационного файла", err)
		panic("Config Error")
	}
}

func GetConfiguration() AppConfiguration {
	return *conf
}