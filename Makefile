.PHONY: run tidy compose-up compose-down db-up db-down

run:
	air

tidy:
	go mod tidy

compose-up:
	docker compose up -d

compose-down:
	docker compose down
db-up:
	migrate -database '$MDSN' -path migrations -verbose up
db-down:
	migrate -database '$MDSN' -path migrations -verbose down