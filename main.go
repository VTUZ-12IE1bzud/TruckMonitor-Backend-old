package main

import (
	"./config"
	"./data"
	"./presentation"
	"fmt"
)

func main() {
	fmt.Println("Старт приложения")
	var conf = config.GetAppConfiguration()
	data.ConnectionDB(conf.DataBaseUrl, conf.DataBaseName, conf.DataBaseUser, conf.DataBasePassword)
	defer data.CloseDB()
	config.GetSchema().Setup()
	config.GetDefaultData().Setup()

	presentation.StartRoute()
}