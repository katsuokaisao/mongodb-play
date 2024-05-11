package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var findByDateRangeCmd = &cobra.Command{
	Use:   "find-by-date-range",
	Short: "Find comments by date range",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		start := "1973-07-26T00:00:00Z"
		st, err := time.Parse(time.RFC3339, start)
		if err != nil {
			log.Fatalf("failed to parse time: %v", err)
		}
		end := "1973-07-28T00:00:00Z"
		ed, err := time.Parse(time.RFC3339, end)
		if err != nil {
			log.Fatalf("failed to parse time: %v", err)
		}
		fmt.Printf("Find by date range: %s ~ %s\n", start, end)

		comments, err := commentRepository.FindByDateRange(st, ed)
		if err != nil {
			log.Fatalf("failed to find comments: %v", err)
		}
		if len(comments) == 0 {
			fmt.Println("comments not found")
		} else {
			if len(comments) > 4 {
				comments = comments[:4]
			}
			printJSON(comments)
		}
	},
}
