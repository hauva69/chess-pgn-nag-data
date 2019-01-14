all: vet lint build

build:
	go build -race

vet:
	go vet

lint:
	golint
