export DATABASE_DRIVER=postgres
export DATABASE_CONNECTION_STRING=postgres://postgres:postgres@localhost/todo_list?sslmode=disable

run:
	make docker-start
	go run ./cmd/main.go

test:
	make docker-start
	go test ./...
	make docker-stop

docker-start:
	@docker-compose up -d && \
	sleep 3

docker-stop:
	@docker-compose down
