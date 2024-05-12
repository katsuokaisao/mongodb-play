package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deleteOneCmd = &cobra.Command{
	Use:   "delete-one",
	Short: "Delete one comment by ID",
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

		err := commentRepository.DeleteOne(_id)
		if err != nil {
			log.Fatalf("failed to delete one comment: %v", err)
		}

		{
			comment, err := commentRepository.FindOneByID(_id)
			if err != nil {
				log.Fatalf("failed to find one comment: %v", err)
			}
			if comment == nil {
				fmt.Println("deleted successfully")
			}
		}
	},
}
