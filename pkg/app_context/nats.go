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
	nc, err := nats.Connect("nats://localhost:4222")

	if err != nil {
		log.Fatal(err)
	}

	return &natsComp{nc: nc}
}

func (c *natsComp) Publish(ctx context.Context, topic string, evt *common.AppEvent) error {
	dataByte, err := json.Marshal(evt.Data)

	if err != nil {
		return errors.WithStack(err)
	}

	return c.nc.Publish(topic, dataByte)
}
