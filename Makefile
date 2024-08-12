.PHONY: up down build restart logs

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