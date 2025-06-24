package myDB

import (
	"log"
	"github.com/angel-omniful/oms/myContext"
	"github.com/omniful/go_commons/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection1 *mongo.Collection
var collection2 *mongo.Collection
var collection3 *mongo.Collection
var collection4 *mongo.Collection
var collection5 *mongo.Collection

func init(){
	ctx := myContext.GetContext()
	//connectionString:= config.GetString(ctx, "mongodb.uri")
	dbname:= config.GetString(ctx, "mongodb.dbname")
	colname1:= config.GetString(ctx, "mongodb.colname1")
	colname2:= config.GetString(ctx, "mongodb.colname2")
	colname3:= config.GetString(ctx, "mongodb.colname3")
	colname4:= config.GetString(ctx, "mongodb.colname4")
	colname5:= config.GetString(ctx, "mongodb.colname5")

	clientOption:=options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB successfully!")

	collection1 = client.Database(dbname).Collection(colname1) //valid orders
	collection2 = client.Database(dbname).Collection(colname2) //invalid orders
	collection3 = client.Database(dbname).Collection(colname3) //webhooks
	collection4 = client.Database(dbname).Collection(colname4) //on hold orders
	collection5 = client.Database(dbname).Collection(colname5) //new orders
	
	if collection1 == nil {
		log.Panicf("Failed to connect to collection %s in database %s", colname1, dbname)
	} else {
		log.Printf("Connected to collection %s in database %s successfully!", colname1, dbname)
	}

	if collection2 == nil {
		log.Panicf("Failed to connect to collection %s in database %s", colname2, dbname)
	} else {
		log.Printf("Connected to collection %s in database %s successfully!", colname2, dbname)
	}

	if collection3 == nil {
		log.Panicf("Failed to connect to collection %s in database %s", colname3, dbname)
	} else {
		log.Printf("Connected to collection %s in database %s successfully!", colname3, dbname)
	}

	if collection4 == nil {
		log.Panicf("Failed to connect to collection %s in database %s", colname4, dbname)
	} else {
		log.Printf("Connected to collection %s in database %s successfully!", colname4, dbname)
	}

	if collection5 == nil {
		log.Panicf("Failed to connect to collection %s in database %s", colname5, dbname)
	} else {
		log.Printf("Connected to collection %s in database %s successfully!", colname5, dbname)
	}

}

func GetOrdersCollection() *mongo.Collection {
	return collection1
}

func GetWebhooksCollection() *mongo.Collection {
	return collection3
}

func GetErrorsCollection() *mongo.Collection {
	return collection2
}

func GetNewOrderCollection() *mongo.Collection {
	return collection5
}

func GetOnholdCollection() *mongo.Collection {
	return collection4
}