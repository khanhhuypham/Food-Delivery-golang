package app_context

import (
	"Food-Delivery/pkg/common"
	"context"
)

type MesssageBroker interface {
	Publish(ctx context.Context, topic string, event *common.AppEvent) error
}
