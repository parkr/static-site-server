PKG=github.com/parkr/static-site-server

all: fmt build test

fmt:
	go fmt $(PKG)/...

build:
	go build $(PKG)

test:
	go test ./...

docker-release:
	docker build -t parkr/static-site-server:$(shell git rev-parse HEAD) .
	docker push parkr/static-site-server:$(shell git rev-parse HEAD)
