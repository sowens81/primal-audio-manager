# Script to validate the Go Application code

Write-Host "Setting working directory"
Set-Location -Path "../"

golangci-lint run --fix

gofmt -w .

goimports -w .

go vet ./...

gosec ./...

Write-Host "Go code validation completed successfully"



