.PHONY: run migrate-up migrate-down

DB_PATH ?= ./out/storage.db

run:
	go run ./cmd/url-shortener

migrate-up:
	migrate -path migrations -database "sqlite://$(DB_PATH)" up
	
migrate-down:
	migrate -path migrations -database "sqlite://$(DB_PATH)" down