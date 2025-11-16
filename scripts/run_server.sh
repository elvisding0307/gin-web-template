set -e

cd src
echo "generating swagger docs"
swag init -g cmd/main.go
go mod tidy
go run cmd/main.go