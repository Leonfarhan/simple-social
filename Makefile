include .env
export

MIGRATION_PATH = cmd/migrate/migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@, $(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@, $(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run ./cmd/migrate/seed

.PHONY: migrate-force
migrate-force:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) force $(filter-out $@,$(MAKECMDGOALS))

.PHONY: dev
dev:
	@colima start && docker compose up -d && air

.PHONY: gen-docs
gen-docs:
	@swag fmt -g main.go -d ./cmd/api,./internal/store
	@swag init -g main.go -d ./cmd/api,./internal/store

# Catch-all target to prevent make from complaining about extra arguments
%:
	@:
