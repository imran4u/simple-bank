
DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable


#Run the postgres container
postgres:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.2-alpine3.20

#Create dabase simaple_bank in running container anem postgres17
createdb: 
	sudo docker exec -it postgres17 createdb --username=root --owner=root simple_bank

# Drop database of container
dropdb: 
	sudo docker exec -it postgres17 dropdb simple_bank

# init migration , -ext sql ( extension), -dir <desination directory>, -seq (sequence)
migrate_init : 
	migrate create -ext sql -dir db/migration -seq init_schema

#to create tables in simple_bank database
migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# To generate sqlc go code from sqlc query ./db/query/ *.sql
sqlc_gen:
	sqlc generate

.PHONY: createdb dropdb postgres migrate_init migrate_up migrate_down sqlc_gen
