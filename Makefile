ROBOT-BIN ?= robot-twitter-search
LDFLAGS ?= -extldflags "-static"

build-api:
	go build -o "$(ROBOT-BIN)" -ldflags '$(LDFLAGS)' cmd/main.go

run-robot:
	swag init -g cmd/maing.go && go run cmd/main.go

test:
	go test ./... -v -coverprofile=coverage.out -count=1
	go tool cover -func=coverage.out

generate:
	# mockgen -source=cmd/main.go -destination=internal/mocks/mock_main.go

