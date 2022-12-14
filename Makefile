postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine3.16
createdb:
	docker exec -it 6c99 postgres14 --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres14 createdb dropdb simple_bank
migrateup:	
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
test:
	go test -v -cover ./...
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown  sqlc test		