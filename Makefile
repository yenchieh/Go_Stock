.PHONY: deps vet test dev build clean buildFront

DOCKER_REPO_URL = jack08300/Go_Stock

deps:
	dep ensure

vet:
	go vet

dev:
	DEBUG=true \
	PORT=8119 \
	gin -p 8120 -a 8119 -x node_modules

build: clean buildFront
	GOOS=linux go build -o ./build/main *.go

clean:
	rm -f view/build/*
	rm -rf build/*
	find . -name '*.test' -delete

push-image: build build-image
	docker tag go_stock $(DOCKER_REPO_URL):latest
	docker push $(DOCKER_REPO_URL):latest

buildFront:
	yarn build

build-image:
	docker build --rm -t go_stock:latest .