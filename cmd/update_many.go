package cmd

import (
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	"github.com/katsuokaisao/mongodb-play/repository"
	"github.com/spf13/cobra"
)

var updateManyCmd = &cobra.Command{
	Use:   "update-many",
	Short: "Update many comments",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		movieID := "573a1392f29313caabcd99e3"
		{
			comments, err := commentRepository.FindByMovieID(movieID)
			if err != nil {
				log.Fatalf("failed to find comments: %v", err)
			}
			if len(comments) == 0 {
				fmt.Println("comments not found")
				return
			}
			if len(comments) > 4 {
				comments = comments[:4]
			}
			printJSON(comments)
		}

		cond := repository.FindCondition{
			MovieID: pointer.ToString(movieID),
		}
		field := repository.UpdateFiled{
			Email: pointer.ToString("update-many2@example.com"),
		}
		if err := commentRepository.UpdateMany(cond, field); err != nil {
			log.Fatalf("failed to update many comments: %v", err)
		}

		{
			comments, err := commentRepository.FindByMovieID(movieID)
			if err != nil {
				log.Fatalf("failed to find comments: %v", err)
			}
			if len(comments) == 0 {
				fmt.Println("comments not found")
				return
			}
			if len(comments) > 4 {
				comments = comments[:4]
			}
			printJSON(comments)
		}
	},
}
