package cmd

import (
	"Food-Delivery/config"
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/db/mysql"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "start consumer",
}

var consumerIncreaseLikeCountRestaurantCmd = &cobra.Command{
	Use:   "increase-like",
	Short: "Start consumer increase like count restaurant",
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := config.LoadConfig("./config/config-local.yml")

		if err != nil {
			log.Println("db connection err: ", err)
			return
		}

		db, err := mysql.MySQLConnection(cfg)
		db = db.Debug()

		if err != nil {
			log.Println("db connection err: ", err)
			return
		}

		nc, err := nats.Connect(cfg.Nat.Url)

		if err != nil {
			log.Fatal("failed to connect nats", err)
		}

		nc.Subscribe(common.EventUserLikeRestaurant, func(msg *nats.Msg) {

			var data rating_dto.CreateDTO
			if err := json.Unmarshal(msg.Data, &data); err != nil {
				log.Println("failed to unmarshal data:", err)
				return
			}

			var str string

			if *data.Like {
				str = "like_count + 1"
			} else {
				str = "like_count - 1"
			}

			var (
				tableName string
				targetID  *int
			)

			switch {
			case data.RestaurantId != nil:
				tableName = model.Restaurant{}.TableName()
				targetID = data.RestaurantId
			case data.ItemId != nil:
				tableName = model.Item{}.TableName()
				targetID = data.ItemId
			default:
				log.Println("No valid target ID found")
				return
			}

			if err := db.Table(tableName).Where("id = ?", *targetID).
				Update("like_count", gorm.Expr(str)).Error; err != nil {
				log.Println("Update like_count failed:", err)
				return
			}

			log.Printf("Update like count success for %s ID: %d", tableName, *targetID)
		})

		// Setup graceful shutdown
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive a signal
		log.Println("Consumer started. Press Ctrl+C to exit...")
		<-signalChan

		log.Println("Shutting down consumer...")

		// Drain connection (process pending messages before closing)
		if err := nc.Drain(); err != nil {
			log.Printf("Error draining NATS connection: %v", err)
		}

		// Close NATS connection
		nc.Close()

		log.Println("Consumer shutdown complete")
	},
}

var consumerDecreaseLikeCountRestaurantCmd = &cobra.Command{
	Use:   "decrease-like",
	Short: "Start consumer decrease like count restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start consumer decrease like count restaurant")
	},
}

func setupConsumerCmd() {
	consumerCmd.AddCommand(consumerDecreaseLikeCountRestaurantCmd)
	consumerCmd.AddCommand(consumerIncreaseLikeCountRestaurantCmd)
}
