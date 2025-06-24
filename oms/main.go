package main

import (
	"context"
	"time"
	myclient "github.com/angel-omniful/oms/client"
	"github.com/angel-omniful/oms/handlers"
	"github.com/angel-omniful/oms/myContext"
	_ "github.com/angel-omniful/oms/myDB"
	"github.com/angel-omniful/oms/routes"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
)

func main(){

	//config init happening in mycontext
	ctx:= myContext.GetContext()
	//client:=myContext.GetHttpClient()
	//logger
	lvl := config.GetString(ctx, "log.level")
	log.SetLevel(lvl)
	log.Infof("Starting OMS on port %d", config.GetInt(ctx, "server.port"))
	log.Println("Connecting to the server...")

	//server
	port:= config.GetString(ctx, "server.port")
	port = ":" + port
	server:= http.InitializeServer(
	port,            
	config.GetDuration(ctx,"server.read_timeout"),      
	config.GetDuration(ctx,"server.write_timeout"),      
	config.GetDuration(ctx,"server.idle_timeout"),       
	false,
	
	)


	//routes
	log.Println("Registering routes...")
	routes.RegisterAllRoutes(server)
	log.Println("Routes registered successfully!")


	//s3
	_, err := myclient.NewS3Client(ctx)
	if err != nil {
		log.Panicf("❌ Failed to initialize S3 client: %v", err)
	}else{
		log.Info("✅ S3 client initialized successfully")
	}

	//upload csv here via api call-
	//right now uploading manually
// curl -X POST http://localhost:8081/api/oms/csv/upload \
//   -H "Content-Type: application/json" \
//   -d '{
//     "local_file_path": "C:/Users/angel/OneDrive/Desktop/OMNIFUL-PROJECT/oms/uploads/test.csv",
//     "key": "uploads/test.csv"
//   }'

	//url:=get-
	// curl -X GET "http://localhost:8081/api/oms/csv/url/test.csv"
	//right now http not working thats why created a temporary function

	// Create request
// request := &http.Request{
// 	Url: "http://localhost:8081/api/oms/csv/test.csv",
// 	Body: map[string]interface{}{
// 	},
// 	Headers: map[string][]string{
// 		"Content-Type": {"application/json"},
// 	},
// 	Timeout: 5 * time.Second, // Optional request-specific timeout
// }

// // For GET request
// 			var response struct {
// 				// Add fields according to the actual JSON response structure
// 				Status  int `json:"status"`
// 				Data    interface{} `json:"data"`
// 				Message string `json:"message"`
// 				Error string `json:"error"`
// 				URL string `json:"url"`
// 			}
// 			resp, err := client.Get(request, &response)
// 			if err!=nil{
// 				log.Errorf("error in fetchinh url for test.csv %v %v",err,resp)
// 			}
// 			url:=response.URL


	url:=handlers.GenerateCsvUrlFunc("test.csv")
	err=handlers.ValidateCSVURL(url)
	//sqs
	sqs, err2 := myclient.NewSQSClient(ctx)
	if err2 != nil {
		log.Panicf("❌ Failed to initialize SQS client: %v", err2)
	}else{
		log.Info("✅ SQS client initialized successfully")
	}

	if err!=nil{
		log.Errorf("url not valid")
	}else{
		log.Info("url validated now uploading")
		data := []byte(url)
		sqs.PublishCreateBulkOrderEvent(ctx,data)
	}

	//sqs consumer will download file
	//it will parse it
	//validate
	//mongo+csv
	log.Info("intialising consumer")
	
	go func(){
	consumer,_:=myclient.NewSQSConsumer(context.Background())

	if consumer==nil{
		log.Info("consumer not created")
	}else{
		log.Info("consumer is:",consumer.Name)
	}
	}()
	

	//kafka
	myclient.InitKafkaProducer(ctx)
	myclient.NewKafkaConsumer(ctx)

	time.Sleep(5 * time.Second) 
	myclient.PublishOrderCreated(ctx)

	//passes mongo ones
	//kafka consumer
	//takes them and checks inventory and register webhooks
	//valid+onhold
//server
	if err := server.StartServer("my-service"); err != nil {
		log.Println("server not connected")
     }else{
		log.Println("server connected successfully")
	}

	select{}

}