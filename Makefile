local:
	@docker-compose -f deploy/local-infrastructure.yaml up --detach

local-stop:
	@docker-compose -f deploy/local-infrastructure.yaml down

run:
	go run ./cmd/not_amazing_amazon

apply:
	kubectl apply \
		-f ./deploy/postgresql.yaml \
		-f ./deploy/minio.yaml \
		-f ./deploy/redis.yaml

push-image: build-image
	docker push indigowar/not_amazing_amazon:latest

build-image: test
	docker build -f ./build/Dockerfile -t indigowar/not_amazing_amazon:latest .

test: lint
	go test ./...

lint: gen
	go vet ./...
	golangci-lint run

gen:
	templ generate ./...
	go generate ./...

deps:
	go mod tidy
	go mod download
