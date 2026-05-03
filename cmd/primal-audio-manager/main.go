package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sowens81/primal-audio-manager/internal/cli"
	"github.com/sowens81/primal-audio-manager/internal/sync"
	"github.com/sowens81/primal-audio-manager/pkg/discogs"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("DISCOGS_TOKEN")
	username := os.Getenv("DISCOGS_USERNAME")

	if token == "" || username == "" {
		log.Fatal("DISCOGS_TOKEN and DISCOGS_USERNAME must be set")
	}

	client := discogs.NewClient(token)

	svc := sync.NewService(client.Collection, username, os.Stdout)

	if err := svc.SyncCollection(); err != nil {
		cli.HandleError(err)
	}

	fmt.Println("✅ sync complete")
}
