.PHONY: up down build restart logs create-migration migrate-up migrate-down

# Up development environment
up:
	docker compose up air -d

# Down development environment
down:
	docker compose down air

# Up with forced build the development environment
build:
	docker compose up air --build -d

# Restart development environment
restart:
	docker compose down air
	docker compose up air -d

# Show logs of development environment
logs:
	docker compose logs -f

# Create a new migration
create-migration:
	go run scripts/create_migration/main.go $(name)

# Run migrations
migrate-up:
	go run scripts/run_migrations/main.go up

# Rollback migrations
migrate-down:
	go run scripts/run_migrations/main.go down