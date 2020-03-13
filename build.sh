set -e

export GOBIN=$(pwd)/bin
go clean *.go
go build *.go
go install *.go
