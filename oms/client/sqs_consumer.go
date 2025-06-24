package client

import (
	"context"
	"os"
	"strconv"
	"time"

	gcConfig "github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/sqs"
	"encoding/csv"
	nethttp "net/http"
	"github.com/angel-omniful/oms/model"
	"github.com/angel-omniful/oms/myContext"
	"github.com/angel-omniful/oms/myDB"
	"github.com/omniful/go_commons/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type MessageHandler struct{}

// Process implements the ISqsMessageHandler interface.
func (h *MessageHandler) Process(ctx context.Context, msgs *[]sqs.Message) error {
	for _, msg := range *msgs {
		log.Info("Received message: %s", string(msg.Value))

		// Step 1: Extract CSV URL from message body
		csvUrl:=string(msg.Value)
		client:=myContext.GetHttpClient()



		resp, err := nethttp.Get(csvUrl)
		if err != nil {
		log.Error("Failed to GET CSV file: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != nethttp.StatusOK {
		log.Error("Failed to fetch file, status code: %d", resp.StatusCode)
		}

		reader := csv.NewReader(resp.Body)
		records, err := reader.ReadAll()
		if err != nil {
		log.Error("Failed to read CSV: %v", err)
		}

		for _, record := range records {
			log.Println("Row:", record)
		}
		log.Info("Parsed %d rows", len(records)-1)

		errors:=myDB.GetErrorsCollection()
		orders:=myDB.GetOrdersCollection()
		// Step 4: Process each record
		for i, row := range records {
			if i == 0 {
				continue // skip header
			}
			quantity, err := strconv.Atoi(row[6])
			if err != nil {
				log.Errorf("invalid quantity: %v", err)
				return err
			}

			createdTime, err := time.Parse(time.RFC3339, row[7])
			if err != nil {
				log.Errorf("invalid date format: %v", err)
				return err
			}

			order:=&model.Order{
				ID:        row[0],
				TenantID:  row[1],
				SellerID:  row[2],
				HubID:     row[3],
				SkuID:     row[4],
				Status:    "on_hold",
				Quantity:  quantity,
				CreatedAt: primitive.NewDateTimeFromTime(createdTime),
	        }

			url:="http://localhost:8080/api/ims/inventory/hubsku/" + order.HubID + "/" + order.SkuID
			
			request := &http.Request{

				Url: url,
				Body: map[string]interface{}{
					},
				Headers: map[string][]string{
					"Content-Type": {"application/json"},
					},
				Timeout: 5 * time.Second, 
			}

			// Define a struct matching the expected inventory API response
			var response struct {
				
				Status  int `json:"status"`
				Data    interface{} `json:"data"`
				Message string `json:"message"`
				Error string `json:"error"`
			}
			resp, err := client.Get(request, &response)

			if err!=nil{
				log.Error("failed to GET inventory: %v %v", err,resp)
			}else if response.Status==200{
				log.Info("suceeded to FETCH inventory")
				res,err:=orders.InsertOne(ctx,order)
				if err!=nil{
					log.Error("error in inserting:",err)
				}else{
					
					log.Info("inserted into valid_order:",res)
				}
				
			}else{
				log.Info("Order doesnt exist in inventory %v", err)
				res,err:=errors.InsertOne(ctx,order)
				if err!=nil{
					log.Error("error in inserting:",err)
				}else{
					
					log.Info("inserted into invalid_order:",res)
				}

				
				
			}

		}
	}
	return nil
}

func NewSQSConsumer(ctx context.Context)(*sqs.Consumer,error){
	os.Setenv("AWS_ACCESS_KEY_ID", gcConfig.GetString(ctx, "aws.sqs.accessKeyId"))
	os.Setenv("AWS_SECRET_ACCESS_KEY", gcConfig.GetString(ctx, "aws.sqs.secretAccessKey"))
	os.Setenv("LOCAL_SQS_ENDPOINT", gcConfig.GetString(ctx, "aws.sqs.endpoint"))
	
	region:= gcConfig.GetString(ctx, "aws.region")
	account:=gcConfig.GetString(ctx, "aws.sqs.account")
	endpoint:=gcConfig.GetString(ctx, "aws.sqs.endpoint")

	queueName:=gcConfig.GetString(ctx,"aws.sqs.queue_name")
	sqsConfig:=sqs.GetSQSConfig(ctx,true,"",region,account,endpoint)
	queue, err := sqs.NewStandardQueue(ctx, queueName, sqsConfig)
	if err != nil {
		log.Error("NewSQSClient: failed to create SQS queue: %v", err)
		return nil, err
	}
  var handler sqs.ISqsMessageHandler = &MessageHandler{}
    consumer, err := sqs.NewConsumer(
        queue,
        1,       // Number of workers
        1,       // Concurrency per worker
        handler,
        10,      // Max messages count
        30,      // Visibility timeout
        false,   // Is async
        false,   // Send batch message
    )
    if err != nil {
        log.Error("Failed to create consumer: %v", err)
		return nil,err
    }

    consumer.Start(ctx)

    // Let the consumer run for a while
    time.Sleep(10 * time.Second)
    consumer.Close()
    return consumer,nil
}


