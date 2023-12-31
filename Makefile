local-postgres:
	docker run --name database -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=root -e POSTGRES_DB="e-wallet" -p 5432:5432 -d postgres:15.4-alpine

create-local-db:
	docker exec -it database createdb --username=root --owner=root "e-wallet"

drop-local-db:
	docker exec -it database dropdb --username=root --owner=root "e-wallet"

migrate-up:
	migrate -path db/migrations -database "postgresql://root:pass@localhost:5432/e-wallet?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgresql://root:pass@localhost:5432/e-wallet?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: local-postgres create-local-db drop-local-db migrate-up migrate-down sqlc test