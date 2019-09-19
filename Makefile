GIT_COMMIT := $(shell git rev-parse --short HEAD)

# builds a Docker image similar to, but not necessarily identical with the one architect would build.
docker-build:
	docker run --rm -v $(shell pwd):/go/src/github.com/giantswarm/health-service \
		-e GOPATH=/go -e GOOS=linux -e GOARCH=amd64 -e CGOENABLED=0 \
		-w /go/src/github.com/giantswarm/health-service \
		quay.io/giantswarm/golang:1.10.4 go build -a -tags netgo -o ./health-service \
		-ldflags "-w -X main.gitCommit=$(GIT_COMMIT) -linkmode 'external' -extldflags '-static'"
	docker build -t quay.io/giantswarm/health-service:latest .
