env-up:
	docker compose up -d

env-down:
	docker compose down

env-cleanup:
	@read -p "Удалить локальную базу данных SQLite? [y/N]: " ans; \
	if [ "$$ans" = "y" ] || [ "$$ans" = "Y" ]; then \
		docker compose down; \
		rm -rf ./out; \
		echo "Папка ./out удалена"; \
	else \
		echo "Отменено"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Ошибка: требуется параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir ./migrations -seq "$(seq)"

migrate-up:
	migrate -path ./migrations -database "sqlite3://./out/storage.db" up

migrate-down:
	migrate -path ./migrations -database "sqlite3://./out/storage.db" down

run:
	docker compose pull && docker compose up -d
