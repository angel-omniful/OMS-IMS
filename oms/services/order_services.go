package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/angel-omniful/oms/model"
	"github.com/angel-omniful/oms/myDB"	
)
var orderCollection *mongo.Collection
var failedOrderCollection *mongo.Collection


func CreateOrder(ctx context.Context, order *model.Order) error {
	orderCollection =myDB.GetOrdersCollection()
	_, err := orderCollection.InsertOne(ctx, order)
	return err
}

func GetOrdersByField(ctx context.Context, field string, value string) ([]model.Order, error) {
	orderCollection =myDB.GetOrdersCollection()
	var orders []model.Order
	cursor, err := orderCollection.Find(ctx, bson.M{field: value})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func CreateFailedOrder(ctx context.Context, failed model.Order) error {
	failedOrderCollection = myDB.GetErrorsCollection()
	_, err := failedOrderCollection.InsertOne(ctx, failed)
	return err
}

func GetAllFailedOrders(ctx context.Context) ([]model.Order, error) {
	failedOrderCollection = myDB.GetErrorsCollection()
	var failed []model.Order
	cursor, err := failedOrderCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &failed); err != nil {
		return nil, err
	}
	return failed, nil
}