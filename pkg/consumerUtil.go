package pkg

import (
	"errors"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer interface {
	ConsumeEvents() (bool, *kafka.Message, error)
	CloseConsumerClient() error
}

type IKafkaConsumer interface {
	CreateConsumer(brokers, consumerGroup, username, password, autoOffsetReset, protocol, mechanisms string, fetchMinBytes, sessionTimeout int, enableAutoCommit bool) (*kafka.Consumer, error)
	SubscribeTopics(consumer *kafka.Consumer, pollTimeOut int) (kafka.Event, error)
}

type ConsumerClient struct {
	consumer     *kafka.Consumer
	consumerImpl IKafkaConsumer
	pollTimeout  int
}

type ConsumerConfigurations struct {
	Brokers, ConsumerGroup, UserName, Password, AutoOffsetReset, Protocol, Mechanisms string
	Topics                                                                            []string
	FetchMinBytes, PollTimeOut, SessionTimeout                                        int
	EnableAutoCommit                                                                  bool
}

type ConsumerUtilService struct {
	ConsumerUtilOperation KafkaConsumer
}

func NewConsumerClient(config *ConsumerConfigurations, consumerImpl IKafkaConsumer) (ConsumerClient, error) {
	consumer, err := consumerImpl.CreateConsumer(config.Brokers, config.ConsumerGroup,
		config.UserName, config.Password, config.AutoOffsetReset, config.Protocol, config.Mechanisms, config.FetchMinBytes, config.SessionTimeout, config.EnableAutoCommit)

	if err != nil {
		log.Println(err.Error())
		return ConsumerClient{
			consumer: nil,
		}, errors.New("kafka consumer client building failed")
	}
	consumer.SubscribeTopics(config.Topics, nil)
	return ConsumerClient{
		consumer:     consumer,
		consumerImpl: consumerImpl,
		pollTimeout:  config.PollTimeOut,
	}, nil

}

func (consumeClient *ConsumerClient) ConsumeEvents() (bool, string, map[string]string, error) {
	ev, _ := consumeClient.consumerImpl.SubscribeTopics(consumeClient.consumer, consumeClient.pollTimeout)
	headers := make(map[string]string)
	switch e := ev.(type) {
	case *kafka.Message:
		for _, v := range e.Headers {
			headers[v.Key] = string(v.Value)
		}
		return true, string(e.Value), headers, nil
	case *kafka.Error:
		consumeErr := fmt.Errorf("error: %v: %v", e.Code(), e)
		if e.Code() == kafka.ErrAllBrokersDown {
			return false, "", nil, consumeErr
		}
		return true, "", nil, consumeErr
	default:
		return true, "InvalidEvent", nil, nil
	}
}

func (consumerClient *ConsumerClient) CloseConsumerClient() error {
	err := consumerClient.consumer.Close()

	if err != nil {
		log.Println("Error in CloseConsumerClient: " + err.Error())
	}
	return nil
}
