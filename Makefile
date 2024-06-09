.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

rundb:
	docker run --name restapi-pg-13.3 -p 25432:5432 -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=dbpass -e POSTGRES_DB=restapi_dev -d postgres:13.3

runtestdb:
	docker run --name restapi-test-pg-13.3 -p 35432:5432 -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=dbpass -e POSTGRES_DB=restapi_test -d postgres:13.3


LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=25432 dbname=restapi_dev user=dbuser password=dbpass sslmode=disable"
LOCAL_MIGRATION_DSN_TEST="host=localhost port=35432 dbname=restapi_test user=dbuser password=dbpass sslmode=disable"

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

create-migration:
	goose -dir ${LOCAL_MIGRATION_DIR} create create_users sql

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

local-test-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} status -v

local-test-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} up -v

local-test-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} down -v