package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var findByNameCmd = &cobra.Command{
	Use:   "find-by-name",
	Short: "Find comments by name",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		name := "Jaqen H'ghar"
		fmt.Printf("Find by name: %s\n", name)
		comments, err := commentRepository.FindByName(name)
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
