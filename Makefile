-include .env
.env: ;
export

MIGRATION_PATH = cmd/migrate/migrations
SWAG_DIRS = ./cmd/api,./internal/store
ARGS = $(filter-out $@,$(MAKECMDGOALS))
ARG_TARGETS = migrate-create migrate-down migrate-force
KNOWN_TARGETS = migrate-create migrate-up migrate-down seed migrate-force dev gen-docs

run = @printf '==> %s\n' "$(1)"; $(2) || { status=$$?; printf 'ERROR: %s failed (exit %s)\n' "$(1)" "$$status" >&2; exit $$status; }

.PHONY: migrate-create
migrate-create:
	$(call run,Creating migration,migrate create -seq -ext sql -dir "$(MIGRATION_PATH)" $(ARGS))

.PHONY: migrate-up
migrate-up:
	$(call run,Running migration up,migrate -path="$(MIGRATION_PATH)" -database="$(DB_ADDR)" up)

.PHONY: migrate-down
migrate-down:
	$(call run,Running migration down,migrate -path="$(MIGRATION_PATH)" -database="$(DB_ADDR)" down $(ARGS))

.PHONY: seed
seed:
	$(call run,Running seed,go run ./cmd/migrate/seed)

.PHONY: migrate-force
migrate-force:
	$(call run,Forcing migration version,migrate -path="$(MIGRATION_PATH)" -database="$(DB_ADDR)" force $(ARGS))

.PHONY: dev
dev:
	$(call run,Starting dev server,colima start && docker compose up -d && air)

.PHONY: gen-docs
gen-docs:
	$(call run,Format Swagger docs,swag fmt -g main.go -d $(SWAG_DIRS))
	$(call run,Generate Swagger docs,swag init -g main.go -d $(SWAG_DIRS))

# Catch-all target to allow extra arguments only for targets that need them.
%:
	@if [ -n "$(filter $(firstword $(MAKECMDGOALS)),$(ARG_TARGETS))" ]; then \
		:; \
	else \
		printf 'ERROR: unknown target "%s".\n' "$@" >&2; \
		printf 'Available targets: %s\n' "$(KNOWN_TARGETS)" >&2; \
		exit 2; \
	fi
