include .env

test:
	@API_KEY=$(API_KEY) go test -v ./...

test-usage:
	@API_KEY=$(API_KEY) go run ./usage/...
	
echo:
	@echo API_KEY=$(API_KEY)
	