run:
	go run ./cmd/app/main.go

goose-status:
	goose -dir ./migrations postgres postgresql://postgres:pass@localhost:6000/postgres?sslmode=disable status

goose-up:
	goose -dir ./migrations postgres postgresql://postgres:pass@localhost:6000/postgres?sslmode=disable up
