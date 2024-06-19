.PHONY: run test local-db lint db/migrate

run:
	air -c .air.toml

test:
	echo hello

local-db:
	docker-compose --env-file ./.env -f docker-compose.yaml -p "productservice" down
	docker-compose --env-file ./.env -f docker-compose.yaml -p "productservice" up -d

lint:
	@(hash golangci-lint 2>/dev/null || \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(go env GOPATH)/bin v1.54.2)
	@golangci-lint run

db/migrate:
	go run ./cmd/migrate

init/db:
	chmod +x initdb.sh
	./initdb.sh

unit-test:
	@mkdir coverage || true
	-go test -p 1 -race -v -coverpkg=./... -coverprofile=coverage/coverage.txt.tmp -count=1  ./...
	@cat coverage/coverage.txt.tmp | grep -v -f .coverageignore > coverage/coverage.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index.html

swagger:
	swag init -g cmd/server/main.go

open-coverage:
	@open coverage/index.html