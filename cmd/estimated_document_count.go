package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var estimatedDocumentCountCmd = &cobra.Command{
	Use:   "estimated-document-count",
	Short: "Get estimated document count",
	Run: func(cmd *cobra.Command, args []string) {
		commentRepository := initMongoDB()

		count, err := commentRepository.EstimatedDocumentCount()
		if err != nil {
			log.Fatalf("failed to get estimated document count: %v", err)
		}
		fmt.Printf("estimated document count: %d\n", count)
	},
}
