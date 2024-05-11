package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/katsuokaisao/mongodb-play/domain"
	"github.com/spf13/cobra"
)

var insertManyCmd = &cobra.Command{
	Use:   "insert-many",
	Short: "Insert many comments",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		comments := []domain.Comment{
			{
				Name:    "Jaime Lannister",
				Email:   "foobar@example.com",
				MovieID: "573a1390f29313caabcd516c",
				Text:    "The things I do for love.",
				Date:    time.Now(),
			},
			{
				Name:    "Cersei Lannister",
				Email:   "foobar@example.com",
				MovieID: "573a1390f29313caabcd516c",
				Text:    "When you play the game of thrones, you win or you die.",
				Date:    time.Now(),
			},
		}

		fmt.Printf("Insert many comments: %s, %s\n", comments[0].Name, comments[1].Name)
		ids, err := commentRepository.InsertMany(comments)
		if err != nil {
			log.Fatalf("failed to insert many comments: %v", err)
		}
		fmt.Printf("Inserted comment IDs: %v\n", ids)
	},
}
