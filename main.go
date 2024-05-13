package main

import (
	"attendance-record/infrastructure/router"
	"log"
)

func main() {
	webServer, err := router.NewWebServerFactory(router.InstanceEcho, "8080")
	if err != nil {
		log.Fatalln(err)
	}

	webServer.Listen()
}
