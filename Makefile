all: vet lint test

build:
	go build -race

vet:
	go vet

lint:
	golint

test: build
	./chess-pgn-nag-data | python -m json.tool > /dev/null
