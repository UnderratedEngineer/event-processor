package app

import (
	"fmt"
	"log"

	"github.com/UnderratedEngineer/event-processor/config"
	"github.com/UnderratedEngineer/event-processor/internal/models"
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

	eventCommonUtils := &models.EventCommonUtils{
		KafkaCommonUtil: &models.KafkaCommonUtils{
			//ConsumerImpl: main.Test,
		},
	}
	fmt.Println(eventCommonUtils)
}
