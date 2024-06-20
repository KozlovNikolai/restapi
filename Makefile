.PHONY: build
build:
	go build -v -o restapiserver ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -func coverage.out
	go tool cover -html coverage.out
	rm coverage.out

coverWithoutMocks:
	go test -short -count=1 -race -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "_mock.go" > cover.out
	rm cover.out.tmp
	go tool cover -func cover.out
	go tool cover -html cover.out
	rm cover.out

coverWithoutMocksBLL:
	# go test ./... -coverpkg='./internal/app/.../bll/...' -coverprofile cover.out.tmp
	go test -short -count=1 -race -coverpkg='./internal/app/...','./model/...' -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "_mock.go" > cover.out
	rm cover.out.tmp
	go tool cover -func cover.out
	go tool cover -html cover.out
	rm cover.out


rundb:
	docker run --name restapi-pg-13.3 -p 25432:5432 -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=dbpass -e POSTGRES_DB=restapi_dev -d postgres:13.3

runtestdb:
	docker run --name restapi-test-pg-13.3 -p 35432:5432 -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=dbpass -e POSTGRES_DB=restapi_test -d postgres:13.3


LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=25432 dbname=restapi_dev user=dbuser password=dbpass sslmode=disable"
LOCAL_MIGRATION_DSN_TEST="host=localhost port=35432 dbname=restapi_test user=dbuser password=dbpass sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

create-migration:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} create create_users sql

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

local-test-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} status -v

local-test-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} up -v

local-test-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} down -v

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml