include .env

default: run

MIGRATION_CMD=docker run -it --rm --network host --volume $(PWD)/internal/infra/database:/db migrate/migrate

# Init API server
.PHONY: run
run:
	docker compose -f docker-compose.dev.yml up -d

# Run importer cli
.PHONY: import-build
import-build:
	docker build -f docker/cli/Dockerfile -t cli .

.PHONY: import
import:
	docker run --rm cli

# add new migration
%:
	@:
.PHONY: migration
migration:
	@if [ "$(word 2,$(MAKECMDGOALS))" = "" ]; then \
		echo "Error: migration name is required; Ex: make migration create_some_table"; \
		exit 1; \
	fi
	@$(MIGRATION_CMD) create -ext sql -dir /db/migrations $(word 2,$(MAKECMDGOALS))

# run migrations up
.PHONY: migrate
migrate:
	@$(MIGRATION_CMD) -path=/db/migrations -database "$(DB_URL)" up

# run migrations down
.PHONY: migrate-down
migrate-down:
	@$(MIGRATION_CMD) -path /db/migrations  -database "$(DB_URL)" down
