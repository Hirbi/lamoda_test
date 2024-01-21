build:
	sudo docker-compose build

up: build
	sudo docker-compose up

migrate:
	goose -dir ./migrations/ up postgres "postgresql://goservice:Go-SeRvIcE482@localhost:5432/goservice?sslmode=disable"
