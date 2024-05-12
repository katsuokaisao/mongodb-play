package cmd

import (
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	"github.com/katsuokaisao/mongodb-play/repository"
	"github.com/spf13/cobra"
)

var updateOneCmd = &cobra.Command{
	Use:   "update-one",
	Short: "Update one comment by ID",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		_id := "664020b5bef7a3d5e852e6c1"

		{
			comment, err := commentRepository.FindOneByID(_id)
			if err != nil {
				log.Fatalf("failed to find one comment: %v", err)
			}
			if comment == nil {
				fmt.Println("comment not found")
				return
			}
			printJSON(comment)
		}

		field := repository.UpdateFiled{
			Email: pointer.ToString("update3@example.com"),
		}
		err := commentRepository.UpdateOne(_id, field)
		if err != nil {
			log.Fatalf("failed to update one comment: %v", err)
		}

		{
			comment, err := commentRepository.FindOneByID(_id)
			if err != nil {
				log.Fatalf("failed to find one comment: %v", err)
			}
			if comment == nil {
				fmt.Println("comment not found")
				return
			}
			printJSON(comment)
		}
	},
}
