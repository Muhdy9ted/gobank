migrateup:
	migrate -path db/migration -database "postgresql://admin:admin%40123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:admin%40123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: migrateup migratedown sqlc test server