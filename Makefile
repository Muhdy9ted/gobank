migrateup:
	migrate -path db/migration -database "postgresql://admin:admin%40123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:admin%40so 123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
.PHONY: migrateup migratedown sqlc test