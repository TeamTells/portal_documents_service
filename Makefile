.PHONY: dc run test lint

dc:
	docker-compose up  --remove-orphans --build

generate:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=openapi/cfg.yaml openapi/openapi.yaml

run:
	go build -o app cmd/main.go

test:
	go test -race ./...

lint:
	golangci-lint run