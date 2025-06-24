package services

import (
	"context"
	"github.com/angel-omniful/oms/model"
	"github.com/angel-omniful/oms/myDB"

)

func RegisterWebhook(ctx context.Context, webhook *model.Webhook) error {
	coll := myDB.GetWebhooksCollection()
	
	_, err := coll.InsertOne(ctx, webhook)
	return err
}

