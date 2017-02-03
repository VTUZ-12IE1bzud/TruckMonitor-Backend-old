package main

import (
	"fmt"
	"TruckMonitor-Backend/service"
)

func main()  {
	fmt.Println("Старт приложения")
	app := service.AppService{}
	app.Run()
}