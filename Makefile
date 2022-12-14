
include .env

all: build

build:
	go build -o person .

migrate:
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" < schema.sql
