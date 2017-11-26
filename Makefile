.PHONY: deps vet test dev build clean buildFront goose-up goose-down

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

migrate:
	docker build --rm -t go_stock:latest .

goose-up:
	cd migrateDB && goose postgres "user=go password=aaaaaa dbname=go_stock sslmode=disable" up && cd ..
goose-down:
	cd migrateDB && goose postgres "user=go password=aaaaaa dbname=go_stock sslmode=disable" down && cd ..

# To import csv to the database
# \copy company_list(symbol,name,last_sale,market_cap,ipo_year,sector,industry,summary_quote) from 'company_list.csv' with csv