package app_context

import (
	"Food-Delivery/pkg/common"
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"log"
)

type natsComp struct {
	nc *nats.Conn
}

func NewNatsComp() *natsComp {
	nc, err := nats.Connect("nats://127.0.0.1:4222")

	if err != nil {
		log.Fatal("unable to conect to nats server: ", err.Error())
	}

	return &natsComp{nc: nc}
}

func (c *natsComp) Publish(ctx context.Context, topic string, event *common.AppEvent) error {
	dataByte, err := json.Marshal(event.Data)

	if err != nil {
		return errors.WithStack(err)
	}

	return c.nc.Publish(topic, dataByte)
}
