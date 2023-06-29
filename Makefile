migrate-db-down:
	@migrate \
		-path ./src/infra/storage/postgres/migrations/ \
		-database "postgresql://postgres:supersecret@localhost:5432?sslmode=disable" \
		down

migrate-db-up:
	@migrate \
		-path ./src/infra/storage/postgres/migrations/ \
		-database "postgresql://postgres:supersecret@localhost:5432?sslmode=disable" \
		up

generate-sql:
	@sqlc generate