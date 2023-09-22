build:
	@echo "building for linux"
	env GOOS=linux CGO_ENABLED=0 go build -o ./bin/shrink-url ./cmd/app
	docker build -t shrink-url .

up:
	docker run -p=4000:80 -d --name shrink-url shrink-url

down:
	-docker stop shrink-url
	-docker rm shrink-url

up_build: down build up
