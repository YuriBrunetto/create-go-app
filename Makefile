build:
	@go build -o bin/create-go-app

run: build
	@./bin/create-go-app
