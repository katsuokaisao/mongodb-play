package cmd

import (
	"log"
	"time"

	"github.com/katsuokaisao/mongodb-play/domain"
	"github.com/spf13/cobra"
)

var replaceOneCmd = &cobra.Command{
	Use:   "replace-one",
	Short: "Replace one comment",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		_id := "66402a408ff70e82665e48b4"

		{
			comment, err := commentRepository.FindOneByID(_id)
			if err != nil {
				log.Fatalf("failed to find one comment: %v", err)
			}
			if comment == nil {
				log.Println("comment not found")
				return
			}
			printJSON(comment)
		}

		comment := domain.Comment{
			ID:      _id,
			Name:    "replaced name",
			Email:   "replace@example.com",
			MovieID: "5f5f3b5b5c6b4b3b4b5b5b5b",
			Text:    "replaced comment",
			Date:    time.Now(),
		}
		if err := commentRepository.ReplaceOne(_id, comment); err != nil {
			log.Fatalf("failed to replace one comment: %v", err)
		}

		{
			comment, err := commentRepository.FindOneByID(_id)
			if err != nil {
				log.Fatalf("failed to find one comment: %v", err)
			}
			if comment == nil {
				log.Println("comment not found")
				return
			}
			printJSON(comment)
		}
	},
}
