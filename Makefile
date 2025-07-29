include .env

test:
	@API_KEY=$(API_KEY) go test -v ./...
	
echo:
	echo API_KEY=$(API_KEY)