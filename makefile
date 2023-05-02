postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

create_db:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank

drop_db:
	docker exec -it postgres12 dropdb simplebank

migrate_create:
	migrate create -dir db/migration -ext sql -seq <schema_name>

migrateup:
	migrate -path db/migration -database postgres://root:secret@127.0.0.1:5432/simplebank?sslmode=disable -verbose up

migratedown:
	migrate -path db/migration -database postgres://root:secret@127.0.0.1:5432/simplebank?sslmode=disable -verbose down

sqlc:
	docker pull kjconroy/sqlc

sqlc_init:
	sqlc init

sqlc_generate:
	docker run --rm -v E:\go\simplebank:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres create_db drop_db migrate_create migrateup migratedown sqlc sqlcinit sqlc_generate test