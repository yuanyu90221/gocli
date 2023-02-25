build:
	GOOS=linux go build -o bin/gocli 

run: build
	./bin/gocli