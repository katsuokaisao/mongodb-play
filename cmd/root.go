package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/katsuokaisao/mongodb-play/mongodb"
	"github.com/katsuokaisao/mongodb-play/repository"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findOneByIDCmd)
	rootCmd.AddCommand(findByNameCmd)
	rootCmd.AddCommand(findByMovieIDCmd)
	rootCmd.AddCommand(findByDateRangeCmd)
	rootCmd.AddCommand(findCmd)
	rootCmd.AddCommand(insertOneCmd)
	rootCmd.AddCommand(insertManyCmd)
	rootCmd.AddCommand(updateOneCmd)
}

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func initMongoDB() repository.CommentRepository {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	cli, err := mongodb.ProvideMongoDBCli(uri)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := cli.Disconnect(context.Background()); err != nil {
			log.Fatalf("failed to disconnect from MongoDB: %v", err)
		}
	}()

	commentRepository := mongodb.NewCommentRepository(cli)

	return commentRepository
}

func printJSON(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
