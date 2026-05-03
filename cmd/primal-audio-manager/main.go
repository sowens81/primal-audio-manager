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
	// Load .env file (non-fatal if missing)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	token := os.Getenv("DISCOGS_TOKEN")
	username := os.Getenv("DISCOGS_USERNAME")

	if token == "" || username == "" {
		log.Fatal("DISCOGS_TOKEN and DISCOGS_USERNAME must be set")
	}

	client := discogs.NewClient(token)

	// 👇 use reusable handler
	err = sync.SyncCollection(client, username)
	if err != nil {
		cli.HandleError(err)
	}

	// err = sync.AddCollectionFolder(client, username, "Hardcore")
	// if err != nil {
	// 	cli.HandleError(err)
	// }

	fmt.Println("✅ sync complete")
}
