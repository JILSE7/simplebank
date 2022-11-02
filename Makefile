createdb:
	docker exec -it myFirstPostgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it  myFirstPostgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/phihdn/simplebank/db/sqlc Store