package myContext

import (
	"os"
	"context"
	"time"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	nethttp "net/http"
)

var ctx context.Context
var htclient *http.Client
func init() {
	//mandatory to call config.init() before using the context

	os.Setenv("CONFIG_SOURCE","local")

	log.Println("Initializing config for oms...")
	err := config.Init(time.Second * 10) // this helps to load the config file (yaml)
	if err != nil {
		log.Panicf("Error while initialising config for oms, err: %v", err)
		panic(err)
	}

	log.Println("Config initialized successfully for oms!")

	log.Println("Creating context for oms...")

	ctx, err = config.TODOContext()
	if err != nil {
		log.Panicf("Failed to create context for oms: %v", err)
	}


	log.Println("Context created successfully for oms!")

	transport := &nethttp.Transport{
	MaxIdleConns:        100,  // Maximum number of idle connections across all hosts
	MaxIdleConnsPerHost: 100,  // Maximum number of idle connections per host
	IdleConnTimeout:     90 * time.Second,  // How long to keep idle connections alive
	}
	
	client, err := http.NewHTTPClient(
	"my-service",           // client service name
	"", // base URL
	transport,
	http.WithTimeout(30 * time.Second), // optional timeout
	)
	htclient=client
	
}

func GetContext() context.Context {
	return ctx
}

func GetHttpClient() *http.Client {
	return htclient
}