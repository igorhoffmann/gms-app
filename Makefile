build:
	docker-compose build GMS-app

run:
	docker-compose up GMS-app


migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up
