# Database
create-migration:
	migrate create -ext sql -dir db/migrations -tz Local $(name)
migrate:
	migrate -database 'postgres://helloWorld:helloWorld@localhost:6432/helloWorld?sslmode=disable' -path db/migrations up
migrate-rollback:
	migrate -database 'postgres://helloWorld:helloWorld@localhost:6432/helloWorld?sslmode=disable' -path db/migrations down $(num)

# App technologies
dependencies-up:
	docker-compose -f deployment/docker-compose.dependencies.yaml up -d
dependencies-down:
	docker-compose -f deployment/docker-compose.dependencies.yaml down
run-app:
	go run main.go