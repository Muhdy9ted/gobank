migrateup:
	migrate -path db/migration -database "postgresql://admin:admin%40123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://admin:admin%40123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://admin:admin%40123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://admin:admin%40123@localhost:5432/go-bank-postgres?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: migrateup migratedown sqlc test server migratedown1 migrateup1