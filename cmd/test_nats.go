package cmd

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"time"
)

var testNatsCmd = &cobra.Command{
	Use:   "test-nats",
	Short: "Test nats",
	Run: func(cmd *cobra.Command, args []string) {
		nc, _ := nats.Connect(nats.DefaultURL)

		// Simple Async Subscriber
		nc.Subscribe("foo", func(m *nats.Msg) {
			fmt.Printf("Received a message: %s\n", string(m.Data))
		})

		// Simple Publisher
		for i := 0; i < 10; i++ {
			nc.Publish("foo", []byte(fmt.Sprintf("Hello World %d", i)))
			time.Sleep(1 * time.Second)
		}

		time.Sleep(20 * time.Second)

	},
}
