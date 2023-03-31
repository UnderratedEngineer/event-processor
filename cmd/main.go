package main

import (
	"github.com/UnderratedEngineer/event-processor/internal/app"
)

type KafkaConsumerImpl struct{}

//this is important to initialize
func InitKafkaConsumer() *KafkaConsumerImpl {
	return &KafkaConsumerImpl{}
}

var Test *KafkaConsumerImpl

func Init() {
	Test = InitKafkaConsumer()
}
func main() {

	app.Start()
}
