package common

type AppEvent struct {
	Topic string
	Data  interface{}
}

// option func
type AppEventOpt func(*AppEvent)

func WithTopic(topic string) AppEventOpt {
	return func(event *AppEvent) {
		event.Topic = topic
	}
}

func WithData(data interface{}) AppEventOpt {
	return func(event *AppEvent) {
		event.Data = data
	}
}

func NewAppEvent(opts ...AppEventOpt) *AppEvent {
	event := &AppEvent{}

	for _, opts := range opts {
		opts(event)
	}

	return event
}
