dev:
	.air -c .air.toml

build:
	go build -o ./dist/main ./internal

start:
	chmod +x ./dist/main
	./dist/main