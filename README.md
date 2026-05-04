# Primal Audio Manager

A Go CLI tool that syncs and manages your Discogs music collection.

## Overview

Primal Audio Manager connects to the Discogs API to sync your vinyl/record collection. It pages through all releases in your collection folders and outputs them to stdout, with a clean sync summary on completion.

Key capabilities:

- Sync all releases across collection folders (paginated)
- Fetch a specific folder by ID
- Add new collection folders
- Structured error handling with actionable Discogs API diagnostics

## Tech Stack

- **Language:** Go 1.25.3
- **API Integration:** Discogs REST API
- **Dependencies:** [`github.com/joho/godotenv`](https://github.com/joho/godotenv)
- **Project Layout:** package-oriented (`pkg/...`, `internal/...`, `cmd/...`)

## Project Structure

``text
.
├─ cmd/
│  └─ primal-audio-manager/   # CLI entrypoint
│     └─ main.go
├─ internal/
│  ├─ errors/                 # Error handling utilities
│  ├─ models/                 # Internal domain models
│  └─ sync/                   # Collection sync service
│     ├─ service.go
│     ├─ types.go
│     └─ service_test.go
├─ pkg/
│  └─ discogs/                # Discogs API client
│     ├─ client.go
│     ├─ errors.go
│     ├─ models/              # Discogs API response models
│     └─ services/
│        └─ collection_service.go
└─ scripts/
   └─ go-validate.ps1
``

## Prerequisites

- Go 1.25.3+
- A [Discogs](https://www.discogs.com) account with a personal access token

## Getting Started

1. Clone the repository:
   ``powershell
   git clone https://github.com/sowens81/primal-audio-manager.git
   cd primal-audio-manager
   ``

2. Install dependencies:
   ``powershell
   go mod tidy
   ``

3. Configure environment variables. You can export them directly or create a `.env` file in the project root:
   ``powershell
   $env:DISCOGS_TOKEN="your-personal-access-token"
   $env:DISCOGS_USERNAME="your-discogs-username"
   ``
   Or create a `.env` file:
   ``env
   DISCOGS_TOKEN=your-personal-access-token
   DISCOGS_USERNAME=your-discogs-username
   ``

4. Run the sync:
   ``powershell
   go run ./cmd/primal-audio-manager
   ``

5. Run tests:
   ``powershell
   go test ./...
   ``

## Sync Service

`internal/sync.Service` orchestrates the collection sync and exposes:

| Method | Description |
|---|---|
| `SyncCollection()` | Pages through all releases in folder 0 (All) and prints them |
| `GetFolderByID(folderID int)` | Fetches and prints metadata for a single folder |
| `AddFolder(name string)` | Creates a new collection folder |
| `GetItemsByFolder(folderID int, pageOpts)` | Fetches and prints releases for a given folder page |

## Discogs Collection API

`pkg/discogs/services.CollectionService` wraps the Discogs API:

- `GetFolders(username string)`
- `GetFolderById(username string, folderID int)`
- `AddFolder(username string, folderName string)`
- `GetItemsByFolder(username string, folderID int, pageOpts)`

## Error Handling

Discogs API errors include:

- HTTP status code
- API error message
- `detail` validation payload
- Raw response body fallback (when JSON parsing fails)

Application errors are handled via `internal/errors.HandleError`.

## Development

Useful commands:

``powershell
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Vet the code
go vet ./...

# Run the validation script
.\scripts\go-validate.ps1
``

## Roadmap

- Improve Discogs error diagnostics
- Expand sync output (CSV / JSON export)
- Add structured logging package
- Add configuration file support

## Contributing

1. Create a branch from `main`
2. Make focused changes
3. Add/adjust tests
4. Open a PR with a clear summary

## License

Add your license here (for example, MIT).