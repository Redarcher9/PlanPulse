hello:
	echo "Hello"

start-client:
	cd ./client/ && yarn start

start-api:
	cd ./api/ && go run main.go