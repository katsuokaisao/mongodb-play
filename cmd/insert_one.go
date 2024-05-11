package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/katsuokaisao/mongodb-play/domain"
	"github.com/spf13/cobra"
)

var insertOneCmd = &cobra.Command{
	Use:   "insert-one",
	Short: "Insert one comment",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		comment := domain.Comment{
			Name:    "Tyrion Lannister",
			Email:   "foobar@example.com",
			MovieID: "573a1390f29313caabcd516c",
			Text:    "I drink and I know things.",
			Date:    time.Now(),
		}
		fmt.Printf("Insert one comment: %s\n", comment.Name)
		id, err := commentRepository.InsertOne(comment)
		if err != nil {
			log.Fatalf("failed to insert one comment: %v", err)
		}
		fmt.Printf("Inserted comment ID: %s\n", id)
	},
}
