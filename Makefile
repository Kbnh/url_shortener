.PHONY: run migrate-up migrate-down

DB_PATH ?= ./out/storage.db

run:
	go run ./cmd/url-shortener

migrate-up:
	$(MAKE) migrate-action action=up

migrate-down:
	$(MAKE) migrate-action action=down

migrate-action:
	@docker compose --env-file .env run --rm url-shortener-migrate \
		-path /migrations \
		-database "sqlite3:///app/out/storage.db" \
		$(action)
