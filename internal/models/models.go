package models

import (
	"database/sql"

	"github.com/UnderratedEngineer/event-processor/pkg"
)

type MongoConnection struct {
	SqlClient   *sql.DB
	MongoConfig string // later replaced with mongo config variable
}

type EventCommonUtils struct {
	KafkaCommonUtil *KafkaCommonUtils
}

type KafkaCommonUtils struct {
	ConsumerImpl pkg.IKafkaConsumer
}
