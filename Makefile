GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

tools: bin/golangci-lint
bin/golangci-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOBIN} v1.37.0

unit-test:
	go test -v ./...

lint:
	$(GOBIN)/golangci-lint run ./...

build:
	cd cmd && go build -o ./bin/cmd && ./bin/cmd

test:
	docker-compose down
	docker-compose up --remove-orphans
	