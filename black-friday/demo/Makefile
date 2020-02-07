APP_VERSION=$(shell git log --pretty=format:'%h' -n 1)
.PHONY: all
all: build push

.PHONY: build
build:
	docker build --compress --rm --build-arg GO_IMAGE=1.12.16-alpine3.10 --build-arg ALPINE_IMAGE=3.10 -t nephe/demo-app:${APP_VERSION} .

.PHONY: push
push:
	docker push nephe/demo-app:${APP_VERSION}
