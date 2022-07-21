build:
	docker-compose build gms-app

run:
	docker-compose up gms-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up
