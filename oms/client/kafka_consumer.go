package client

import (
	"context"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/angel-omniful/oms/model"
	"github.com/angel-omniful/oms/myContext"
	"github.com/angel-omniful/oms/myDB"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/log"
	"time"
)

// Implement message handler
type KafkaMessageHandler struct{}



func (h KafkaMessageHandler) Process(ctx context.Context, msgs *pubsub.Message) error{
	
	//collections i need to work with
    onhold:=myDB.GetOnholdCollection()
	neworder:=myDB.GetNewOrderCollection()
	var order model.Order
	client:=myContext.GetHttpClient()
	
	//working with ip parameter
    if err := json.Unmarshal(msgs.Value, &order); err != nil {
        log.Errorf("❌ Failed to unmarshal order: %v", err)
        return err
    }
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Errorf("❌ Failed to marshal order: %v", err)
		return err
	}
	log.Infof("✅ Processed order: %+v", order)
	webh:=myDB.GetWebhooksCollection()
		var webhook model.Webhook
		filter := bson.M{"tenant_id": order.TenantID, "active": true}
		err = webh.FindOne(ctx, filter).Decode(&webhook)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Errorf("no tenant present")
			}
			log.Errorf("some other error with fetching webhook")
		}

	//Post req for updating inventory as per availability
	url:="http://localhost:8080/api/ims/inventory/check"
	request := &http.Request{
				Url: url,
				Body: orderJSON,
				Headers: map[string][]string{
					"Content-Type": {"application/json"},
					},
				Timeout: 5 * time.Second,
			}
			var response struct {
				Status  int `json:"status"`
				Data    interface{} `json:"data"`
				Message string `json:"message"`
				Error string `json:"error"`
			}
			resp, err := client.Post(request, &response)
			if err!=nil{
				log.Infof("error in req %v %v",err,resp)
			}

	webhookurl:=webhook.URL
    if response.Status == 200{
			order.Status="new_order"
			orderJSON,_:= json.Marshal(order)
            res,err:=neworder.InsertOne(ctx,order)
				if err!=nil{
					log.Error("error in inserting:",err)
				}else{
					
					log.Info("inserted into new_order:",res)
				}
				request2 := &http.Request{
				Url: webhookurl,
				Body: orderJSON,
				Headers: map[string][]string{
					"Content-Type": {"application/json"},
					},
				Timeout: 5 * time.Second,
			}
			resp2, _:= client.Post(request2, &response)

			if response.Status==200{
				log.Info("webhook registered for this order")
			}else{
				log.Infof("unable to register webhook %v",resp2)
			}
    }else{
		
			order.Status="on_hold"
			orderJSON,_:= json.Marshal(order)
       		 res,err:=onhold.InsertOne(ctx,order)
				if err!=nil{
					log.Error("error in inserting:",err)
				}else{
					
					log.Info("inserted into on_hold:",res)
				}
				request2 := &http.Request{

				Url: webhookurl,
				Body: orderJSON,
				Headers: map[string][]string{
					"Content-Type": {"application/json"},
					},
				Timeout: 5 * time.Second,
				}
			resp2, _:= client.Post(request2, &response)

			if response.Status==200{
				log.Info("webhook registered for this order")
			}else{
				log.Infof("unable to register webhook %v",resp2)
			}
    }
	
	log.Infof("✅ Processed order: %+v", order)
	
	return nil
}



func NewKafkaConsumer(ctx context.Context) {
   	brokers := config.GetStringSlice(ctx, "kafka.brokers")
	version := config.GetString(ctx, "kafka.version")
	topic:=config.GetString(ctx, "kafka.topic")
	if version == "" {
		log.Panicf("Kafka version is missing in config")
	}
    consumer := kafka.NewConsumer(
        kafka.WithBrokers(brokers),
        kafka.WithConsumerGroup("oms-group"),
        kafka.WithClientID("oms-consumer"),
        kafka.WithKafkaVersion(version),
    )

	
   //defer consumer.Close()

    // Set NewRelic interceptor for monitoring
    //consumer.SetInterceptor(interceptor.NewRelicInterceptor())
	log.Infof("✅ Subscribing to Kafka topic: %s with group oms-group", topic)

    // Register message handler for topic
	var handler pubsub.IPubSubMessageHandler = KafkaMessageHandler{}

	consumer.RegisterHandler(topic, handler)

    // Start consuming messages
    go consumer.Subscribe(ctx)
}
