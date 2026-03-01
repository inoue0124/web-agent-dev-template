.PHONY: migrate-up migrate-down migrate-create

MIGRATE=migrate
DATABASE_URL ?= postgresql://postgres:postgres@localhost:5432/app_dev?sslmode=disable

migrate-up:
	$(MIGRATE) -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	$(MIGRATE) -path migrations -database "$(DATABASE_URL)" down 1

migrate-create:
	$(MIGRATE) create -ext sql -dir migrations -seq $(name)
