package main

import (
	"attendance-record/infrastructure/log"
	"attendance-record/infrastructure/router"
	"attendance-record/model"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()

	webServer := router.NewWebServerFactory(
		router.InstanceEcho,
		log.InstanceZapLogger,
		"8080",
	)
	webServer.Listen()
}
