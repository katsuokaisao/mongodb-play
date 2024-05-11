package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var findOneByIDCmd = &cobra.Command{
	Use:   "find-one-by-id",
	Short: "Find one comment by ID",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()
		_id := "5a9427648b0beebeb6957a21"
		fmt.Printf("Find by ID: %s\n", _id)
		comment, err := commentRepository.FindOneByID(_id)
		if err != nil {
			log.Fatalf("failed to find one comment: %v", err)
		}
		if comment == nil {
			fmt.Println("comment not found")
		} else {
			printJSON(comment)
		}
	},
}
