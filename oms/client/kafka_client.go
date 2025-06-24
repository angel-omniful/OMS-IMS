package client

import (
	"context"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/pubsub"
	"github.com/angel-omniful/oms/model"
	"github.com/angel-omniful/oms/myDB"
	"go.mongodb.org/mongo-driver/bson"
)

var kafkaLogger = log.DefaultLogger()

var producer *kafka.ProducerClient

func InitKafkaProducer(ctx context.Context){
	brokers := config.GetStringSlice(ctx, "kafka.brokers")
	version := config.GetString(ctx, "kafka.version")

	if version == "" {
		log.Panicf("Kafka version is missing in config")
	}

	producer = kafka.NewProducer(
		kafka.WithBrokers(brokers),
		kafka.WithClientID("oms-producer"),
		kafka.WithKafkaVersion(version), 
	)

	kafkaLogger.Infof("Kafka producer initialized with brokers: %v, version: %s", brokers, version)

	
	
}

func PublishOrderCreated(ctx context.Context) {

	kafkaLogger.Infof("Producer topic: %s", config.GetString(ctx, "kafka.topic"))
	orders:=myDB.GetOrdersCollection()

	filter := bson.M{}

	cursor, err := orders.Find(ctx, filter)
    if err != nil {
        log.Errorf("Find failed: %v", err)
    }
    defer cursor.Close(ctx)

	

	for cursor.Next(ctx){
		var order model.Order

		 if err := cursor.Decode(&order); err != nil {
            log.Printf("Failed to decode document: %v", err)
            continue
		 }

		 	payload, err := pubsub.NewEventInBytes(order)
			if err != nil {
			kafkaLogger.Errorf(" Failed to marshal OrderCreated: %v", err)
			return
			}

			msg := &pubsub.Message{
				Topic: config.GetString(ctx, "kafka.topic"),
				Key:   order.ID,
				Value: payload,
			}
			kafkaLogger.Infof("About to publish to Kafka: topic=%s, key=%s, payload=%s", 
			msg.Topic, msg.Key, string(msg.Value))

			if err := producer.Publish(ctx, msg); err != nil {
				kafkaLogger.Errorf("Kafka publish error: %v", err)
			} else {
				kafkaLogger.Infof("Published order.created for OrderID: %s", order.ID)
			}

	}
	
}
