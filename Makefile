.PHONY test:
test:
	@docker compose -f docker/docker-compose-test.yml up -d

.PHONY: up
up:
	@docker compose -f docker/docker-compose-postgres.yml up -d postgres

.PHONY: down
down:
	@docker compose -f docker/docker-compose-postgres.yml down