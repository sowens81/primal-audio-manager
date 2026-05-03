# Primal Audio Manager

Go-based audio collection manager with Discogs integration.

## Overview

Primal Audio Manager provides services and models for working with Discogs collection data, including:

- Fetching collection folders
- Fetching a folder by ID
- Adding folders
- Fetching releases in a folder

## Tech Stack

- **Language:** Go
- **API Integration:** Discogs
- **Project Layout:** package-oriented (`pkg/...`, `internal/...`)

## Project Structure

```text
.
├─ internal/
│  └─ models/
├─ pkg/
│  └─ discogs/
│     ├─ models/
│     └─ services/
│        └─ collection_service.go
└─ README.md
```

## Prerequisites

- Go 1.22+ (or your repo’s required Go version)
- Discogs account + API credentials (if calling live API)

## Getting Started

1. Clone the repository:
   ```powershell
   git clone https://github.com/sowens81/primal-audio-manager.git
   cd primal-audio-manager
   ```

2. Install dependencies:
   ```powershell
   go mod tidy
   ```

3. Configure environment variables (example):
   ```powershell
   $env:DISCOGS_TOKEN="your-token"
   $env:DISCOGS_USER_AGENT="PrimalAudioManager/1.0"
   ```

4. Run tests:
   ```powershell
   go test ./...
   ```

## Example Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/sowens81/primal-audio-manager/pkg/discogs/services"
)

func main() {
    // Create your Discogs API client (implementation-specific)
    // client := discogs.NewClient(...)

    // collectionSvc := services.NewCollectionService(client)
    // folders, err := collectionSvc.GetFolders("your-username")
    // if err != nil {
    // 	log.Fatal(err)
    // }

    fmt.Println("Initialize service and call methods from pkg/discogs/services")
    log.Println("See collection_service.go for available methods")
}
```

## API Services

`CollectionService` currently exposes methods similar to:

- `GetFolders(username string)`
- `GetFolderById(username string, folderID int)`
- `AddFolder(username string, folderName string)`
- `GetFolderReleases(username string, folderID int)`

## Error Handling

When integrating with Discogs, include:

- HTTP status code
- API error message
- `detail` validation payload
- Raw response body fallback (when JSON parse fails)

This makes validation failures actionable.

## Development

Useful commands:

```powershell
go test ./...
go test -v ./...
go vet ./...
```

## Roadmap

- Improve Discogs error diagnostics
- Add unit tests for service layer
- Add CLI or API entrypoint
- Add configuration and logging package

## Contributing

1. Create a branch from `main`
2. Make focused changes
3. Add/adjust tests
4. Open a PR with a clear summary

## License

Add your license here (for example, MIT).