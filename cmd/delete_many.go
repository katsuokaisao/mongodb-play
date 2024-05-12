package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/katsuokaisao/mongodb-play/repository"
	"github.com/spf13/cobra"
)

var deleteManyCmd = &cobra.Command{
	Use:   "delete-many",
	Short: "Delete many comments",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		start := "1973-07-25T00:00:00Z"
		st, err := time.Parse(time.RFC3339, start)
		if err != nil {
			log.Fatalf("failed to parse time: %v", err)
		}
		end := "1973-07-26T00:00:00Z"
		ed, err := time.Parse(time.RFC3339, end)
		if err != nil {
			log.Fatalf("failed to parse time: %v", err)
		}
		cond := repository.FindCondition{
			Start: pointer.ToTime(st),
			End:   pointer.ToTime(ed),
		}
		{
			comments, err := commentRepository.Find(cond)
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

		if err := commentRepository.DeleteMany(cond); err != nil {
			log.Fatalf("failed to delete many comments: %v", err)
		}

		{
			comments, err := commentRepository.Find(cond)
			if err != nil {
				log.Fatalf("failed to find comments: %v", err)
			}
			if len(comments) == 0 {
				fmt.Println("deleted successfully")
			}
		}
	},
}
