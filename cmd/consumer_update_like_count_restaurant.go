package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "start consumer",
}

var consumerIncreaseLikeCountRestaurantCmd = &cobra.Command{
	Use:   "increase-like",
	Short: "Start consumer increase like count restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start consumer increase like count restaurant")
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
