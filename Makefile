
run: 
	go run cmd/app/main.go

run_db_flag: 
	go run cmd/app/main.go --database="Postgres"
	
migrate:
	migrate -path migrations/ -database "postgres://postgres:4650@localhost/database?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations/ -database "postgres://postgres:4650@localhost/database?sslmode=disable" -verbose down

