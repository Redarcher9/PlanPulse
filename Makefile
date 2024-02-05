hello:
	echo "Hello"

start-client:
	cd ./client/ && yarn start

start-api:
	cd ./server/ && go run cmd/main.go