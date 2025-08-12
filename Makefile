include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-reset
migrate-reset:
	@migrate --path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) drop -f

.PHONY: migrate-reset-up
migrate-reset-up: migrate-reset migrate-up

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go


.PHONY: gen-docs
gen-docs:
	@swag init -g main.go -d cmd/api,internal/store,internal/env,internal/db -o docs && swag fmt