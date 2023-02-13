dev:
	docker-compose -f docker-compose.local.yaml up -d && docker logs -f -n 10 web

build:
	go build -o ./dist/main ./internal

start:
	chmod +x ./dist/main
	./dist/main