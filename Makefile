include .env

test:
	@go test -v .

# Don't use this.
test-e2e:
	@API_KEY=$(API_KEY) go test -v ./e2e/...
	
echo:
	@echo API_KEY=$(API_KEY)
	