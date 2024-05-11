package main

import (
	"log"

	"github.com/katsuokaisao/mongodb-play/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
