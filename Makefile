local-start:
	@docker-compose -f deploy/local-infrastructure.yaml up --detach

local-stop:
	@docker-compose -f deploy/local-infrastructure.yaml down

run: gen
	go run ./cmd/not_amazing_amazon

push-image: test
	docker build -f ./build/Dockerfile -t indigowar/not_amazing_amazon:latest .
	docker push indigowar/not_amazing_amazon:latest

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
