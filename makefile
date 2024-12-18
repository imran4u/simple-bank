
DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable


#Run the postgres container
postgres:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.2-alpine3.20

#Restrat if container is there
postgres-start:
	sudo docker start postgres17 

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
migrate_up1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
migrate_down1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

# To generate sqlc go code from sqlc query ./db/query/ *.sql
sqlc_gen:
	sqlc generate

# command for unit test ./... ( to run in multiple packages)
# ... : This is a Go-specific wildcard that means all subdirectories recursively
# -count=1 : to disable go test cache.
test:
	# go test -v -cover -count=1 ./...

#start server
run: 
	go run main.go

#Generate db mock, at last source folder and interface Name.
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/imran4u/simple-bank/db/sqlc Store

.PHONY: createdb dropdb postgres migrate_init migrate_up migrate_down migrate_up1 migrate_down1 sqlc_gen test postgres-start run mock
