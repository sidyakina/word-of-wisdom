build-all: build-server build-client

build-server:
	docker build -t word-of-wisdom-server -f ./build/server/Dockerfile .
build-client:
	docker build -t word-of-wisdom-client -f ./build/client/Dockerfile .

start-all: start-server start-client

start-server:
	echo "starting word-of-wisdom-server"
	docker run --rm -d --env-file server.env --network host --name word-of-wisdom-server word-of-wisdom-server:latest
	echo "word-of-wisdom-server started"

start-client:
	echo "starting word-of-wisdom-client"
	docker run --rm --env-file client.env --network host --name word-of-wisdom-client word-of-wisdom-client:latest

stop-server:
	docker rm -f word-of-wisdom-server

linter:
	sudo docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.46.2 golangci-lint run -v