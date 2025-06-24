package myContext

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/omniful/go_commons/config"
)

var ctx context.Context

func init() {
	os.Setenv("CONFIG_SOURCE","local")
	//mandatory to call config.init() before using the context
	log.Println("Initializing config...")
	err := config.Init(time.Second * 10) 
	if err != nil {
		log.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}

	log.Println("Config initialized successfully!")

	log.Println("Creating context...")

	ctx, err = config.TODOContext()
	if err != nil {
		log.Panicf("Failed to create context: %v", err)
	}

	log.Println("Context created successfully!")
}

func GetContext() context.Context {
	return ctx
}