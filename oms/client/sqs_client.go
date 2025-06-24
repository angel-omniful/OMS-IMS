package client

import (
	"os"
	"context"

	gcConfig "github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/sqs"
)

// SQSClient wraps the GoCommons SQS Publisher for CreateBulkOrder
type SQSClient struct {
	Publisher *sqs.Publisher
}

// NewSQSClient initializes the SQS publisher using config
func NewSQSClient(ctx context.Context) (*SQSClient, error) {
	os.Setenv("AWS_ACCESS_KEY_ID", gcConfig.GetString(ctx, "aws.sqs.accessKeyId"))
	os.Setenv("AWS_SECRET_ACCESS_KEY", gcConfig.GetString(ctx, "aws.sqs.secretAccessKey"))
	os.Setenv("LOCAL_SQS_ENDPOINT", gcConfig.GetString(ctx, "aws.sqs.endpoint"))

	region:= gcConfig.GetString(ctx, "aws.region")
	account:=gcConfig.GetString(ctx, "aws.sqs.account")
	endpoint:=gcConfig.GetString(ctx, "aws.sqs.endpoint")


	queueName:=gcConfig.GetString(ctx,"aws.sqs.queue_name")
	sqsConfig:=sqs.GetSQSConfig(ctx,true,"",region,account,endpoint)

	// Initialize queue
	queue, err := sqs.NewStandardQueue(ctx, queueName, sqsConfig)
	if err != nil {
		log.DefaultLogger().Errorf("NewSQSClient: failed to create SQS queue: %v", err)
		return nil, err
	}

	// Create publisher
	publisher := sqs.NewPublisher(queue)

	log.DefaultLogger().Infof("SQS Publisher initialized for queue: %s", queueName)

	return &SQSClient{
		Publisher: publisher,
	}, nil
}

// PublishCreateBulkOrderEvent sends a message payload to the SQS queue
func (c *SQSClient) PublishCreateBulkOrderEvent(ctx context.Context, payload []byte) error {
	msg := &sqs.Message{
		Value: payload,
	}

	if err := c.Publisher.Publish(ctx, msg); err != nil {
		log.DefaultLogger().Errorf("SQS publish failed: %v", err)
		return err
	}

	log.DefaultLogger().Infof("SQS message published successfully")
	return nil
}







































// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"
//     "github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/angel-omniful/oms/myContext"
// 	"github.com/omniful/go_commons/config"
// 	"github.com/omniful/go_commons/sqs"
// )

// type MessageHandler struct{}



// // Implement the required Process method for ISqsMessageHandler interface
// func (h *MessageHandler) Process(ctx context.Context, msgs *[]sqs.Message) error {
//     for _, msg := range *msgs {
//         fmt.Println("Processing message:", string(msg.Value))
//     }
//     return nil
// }

// //var publisher *sqs.Publisher
// func InitialiseAws() {
//     // Initialize AWS session
//     ctx:=myContext.GetContext()
//     log.Println("Initializing AWS session...")
//     _, err := session.NewSession(&aws.Config{
//         Region: aws.String(config.GetString(ctx,"aws.region")),
//         Credentials: credentials.NewStaticCredentials("test","test",""), // Use environment variables for credentials
//         Endpoint: aws.String(config.GetString(ctx,"aws.endpoint")), // Optional, specify if using a custom endpoint
//     })
//     if err != nil {
//         log.Fatalf("Failed to create AWS session: %v", err)
//     }
//     log.Printf("AWS session initialized successfully")
    
//     sqsCfg := &sqs.Config{
//         Account:  config.GetString(ctx, "aws.sqs.account"),
//         Endpoint: config.GetString(ctx, "aws.endpoint"),
//         Region:   config.GetString(ctx, "aws.region"),
//     }

//     //creating a queue
//     queue_name:=config.GetString(ctx,"aws.sqs.queue_name")
//     // Initialize Queue
//     queue,err := sqs.NewStandardQueue(ctx,queue_name,sqsCfg)  // Initialize Queue object with sess and queueURL
//     if err!=nil{
//         log.Println("queue not initialized")
//         log.Println(err)
//     }else{
//         log.Println("queue initialized")
//     }
//     // Set up publisher
//     publisher := sqs.NewPublisher(queue)
//     message := &sqs.Message{
//         Value: []byte("Hello SQS!"),
//     }

   
//     if err := publisher.Publish(ctx, message); err != nil {
//         log.Fatalf("Failed to publish message: %v", err)
//     }

//     // Set up consumer
//     var handler sqs.ISqsMessageHandler = &MessageHandler{}
//     consumer, err := sqs.NewConsumer(
//         queue,
//         1,       // Number of workers
//         1,       // Concurrency per worker
//         handler,
//         10,      // Max messages count
//         30,      // Visibility timeout
//         false,   // Is async
//         false,   // Send batch message
//     )
//     if err != nil {
//         log.Fatalf("Failed to create consumer: %v", err)
//     }

//     consumer.Start(ctx)

//     // Let the consumer run for a while
//     time.Sleep(10 * time.Second)
//     consumer.Close()
// }




