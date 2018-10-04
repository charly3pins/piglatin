.PHONY: all docker_test docker_coverage deps build

BIN_NAME=piglatin

all: test docker_test docker_coverage deps build

test:
	mkdir -p cover
	go test  -coverprofile=cover/coverage.out
	go tool cover -html=cover/coverage.out -o cover/index.html

docker_test: 
	docker run --rm -v "${PWD}:/go/src/github.com/charly3pins/piglatin/"  -w "/go/src/github.com/charly3pins/piglatin/" golang:1.10 make deps test

docker_coverage:
	make docker_test
	docker run --rm -v "${PWD}/cover/.:/usr/share/nginx/html" -p 8081:80 nginx

deps:
	go get github.com/charly3pins/piglatin

build:
	cd cmd/server; go build -o ${BIN_NAME}
	@echo "You can now use ./${BIN_NAME}"
