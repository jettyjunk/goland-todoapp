include .env
export 

export PROJECT_ROOT=$(shell pwd)


env-up:
	@docker compose up -d todoapp-postgres port-forwarder
	@sudo chmod -R 777 out/pgdata

env-down:
	@docker compose down todoapp-postgres port-forwarder

env-cleanup:
	@read -p "Clear volumes file? DANGER lose file. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		sudo rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "file env clear"; \
	else \
		echo "file env cancel"; \
	fi 


env-port-forward:
	@docker compose up -d  port-forwarder

env-port-forward-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "no parameter seq: (make migration-create seq=name migrations)"; \
		exit 1; \
	fi; \
	docker compose run --rm --user $$(id -u):$$(id -g) \
		todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down 

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "no parameter action: (make migration-action action=up, down)"; \
		exit 1; \
	fi; \
	docker compose run --rm --user $$(id -u):$$(id -g) todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"



migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "no parameter version: (make migration-force version=1)"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		force $(version)


logs-cleanup:
	@read -p "Clear logs file? DANGER lose file. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		sudo rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "file logs clear"; \
	else \
		echo "file logs clear cancel"; \
	fi 


todoapp-run:
	@export LOGGER_FOLDER=$(PROJECT_ROOT)/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run $(PROJECT_ROOT)/cmd/todoapp/main.go