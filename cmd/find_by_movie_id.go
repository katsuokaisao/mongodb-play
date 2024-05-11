package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var findByMovieIDCmd = &cobra.Command{
	Use:   "find-by-movie-id",
	Short: "Find comments by movie ID",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		movieID := "573a1390f29313caabcd516c"
		fmt.Printf("Find by movie ID: %s\n", movieID)
		comments, err := commentRepository.FindByMovieID(movieID)
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
