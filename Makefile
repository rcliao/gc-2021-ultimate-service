SHELL := /bin/bash

run:
	go run ./app/service/sales-api/

# $(git rev-parse --short HEAD) to get current git sha
VERSION := 1.0

all: sales-api

sales-api:
	docker build \
		-f zarf/docker/Dockerfile.sales-api \
		-t sales-api-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
