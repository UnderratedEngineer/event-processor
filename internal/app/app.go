package app

import (
	"log"

	"github.com/UnderratedEngineer/event-processor/config"
)

func Start() {

	log.Println("Starting the alert event service !!")

	appconfig, configerr := config.Get()

	if configerr != nil {
		log.Println("error occurred while loading config ")
	} else {
		log.Println("config loaded. ")
		log.Println(appconfig)
	}

}
