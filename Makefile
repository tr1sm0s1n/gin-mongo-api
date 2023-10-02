BIN_LOC=bin/gin-mongo-api

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

enter:
	docker exec -it gin-mongo mongosh -u root -p rootpw

compile:
	go build -o ${BIN_LOC} main.go

start:
	./${BIN_LOC}

run:
	go run .

air:
	air
