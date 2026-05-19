.PHONY: run migrate-up migrate-down

DB_PATH ?= out/url-shortener.db

run: migrate-up
	go run ./cmd/url-shortener

migrate-up:
	migrate -path migrations -database "sqlite3://$(DB_PATH)" up
	
migrate-down:
	migrate -path migrations -database "sqlite3://$(DB_PATH)" down