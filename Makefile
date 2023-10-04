postgres :
	sudo docker run --name postgres16 -p 5432:5432 -e  POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
createdb :
	sudo docker exec -it postgres16 createdb --username=postgres --owner=postgres simple_bank 

dropdb :
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	sudo migrate -path db/migration -database "postgres://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

sqlc:
	 sqlc generate

.PHONY : postgres createdb dropdb 