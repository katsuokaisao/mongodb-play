package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/katsuokaisao/mongodb-play/domain"
	"github.com/spf13/cobra"
)

var updateOneCmd = &cobra.Command{
	Use:   "update-one",
	Short: "Update one comment by ID",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		_id := "5a9427648b0beebeb6957a21"
		fmt.Printf("Update by ID: %s\n", _id)
		comment := domain.Comment{
			Name:    "Tyrion Lannister",
			Email:   "update@example.com",
			MovieID: "573a1390f29313caabcd516c",
			Text:    "I drink and I know things.",
			Date:    time.Now(),
		}
		err := commentRepository.UpdateOne(_id, &comment)
		if err != nil {
			log.Fatalf("failed to update one comment: %v", err)
		}
		fmt.Println("updated")
	},
}
