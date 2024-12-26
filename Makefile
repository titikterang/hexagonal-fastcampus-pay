#!/bin/bash
export GOPRIVATE=github.com/titikterang

API_PROTO_FILES=$(shell find lib/protos/$(MODULE) -name *.proto -not -path '*/vendor/*')
API_PROTO_CLIENT=$(shell find lib/protos/openapiv2/$(MODULE) -name *.json -not -path '*/vendor/*')

init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest

.PHONY: print
print:
	@echo $(API_PROTO_FILES)

print-client:
	@echo $(API_PROTO_CLIENT)

.PHONY: generate
generate:
	protoc -I . --proto_path=. \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
	       --openapiv2_out ./lib/protos/openapiv2 --openapiv2_opt logtostderr=true \
			$(API_PROTO_FILES)

generate-js-client:
	openapi-generator generate -g javascript --additional-properties=usePromises=true -i $(API_PROTO_CLIENT) -o ./protos/client/js/$(MODULE)

lint:
	@echo "install golangci-lint from here https://golangci-lint.run/usage/install/ ...."
	@golangci-lint run ./... > .golint.txt

build-membership:
	@CGO_ENABLED=0  go build -o ./bin/membership ./cmd/membership/*

build-banking:
	@CGO_ENABLED=0  go build -o ./bin/banking_api ./cmd/banking/api/*
	@CGO_ENABLED=0  go build -o ./bin/banking_consumer ./cmd/banking/consumer/*

build-money:
	@CGO_ENABLED=0  go build -o ./bin/money_api ./cmd/money/api/*
	@CGO_ENABLED=0  go build -o ./bin/money_consumer ./cmd/money/consumer/*

build-transfer:
	@CGO_ENABLED=0  go build -o ./bin/transfer_api ./cmd/transfer/api/*
	@CGO_ENABLED=0  go build -o ./bin/transfer_consumer ./cmd/transfer/consumer/*

build-payment:
	@CGO_ENABLED=0  go build -o ./bin/payment_api ./cmd/payment/api/*
	@CGO_ENABLED=0  go build -o ./bin/payment_consumer ./cmd/payment/consumer/*

run-mongodb:
	@docker-compose -f docker-compose.mongodb.yaml up

run-membership: build-membership
	@echo "RUN Membership DB & Backend Service ..."
	@docker-compose -f docker-compose.membership.yaml up --force-recreate

run-banking: build-banking
	@echo "RUN Banking DB & Backend Service ..."
	@docker-compose -f docker-compose.banking.yaml up --force-recreate

run-money: build-money
	@echo "RUN Money DB & Backend Service ..."
	@docker-compose -f docker-compose.money.yaml up --force-recreate

run-krakend:
	@docker run --rm --name krakend -p 18080:8080 \
		-v /etc/localtime:/etc/localtime:ro \
		--add-host=host.docker.internal:host-gateway \
		-v ./files/config:/etc/krakend/ \
		-v ./files/jwks:/data/jwks/ devopsfaith/krakend run \
		--config /etc/krakend/krakend.json

run-vault-local:
	@vault server -dev -dev-root-token-id="supersecret"

add-vault-token:
	@echo "usage : POLICY={your_new_policy_name} make add-vault-token"
	@vault token create -policy=$(POLICY) -address=http://localhost:8200