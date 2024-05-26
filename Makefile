docker-up:
	docker-compose up -d

build-app:
	cd cmd/ordersystem && go build -o app -v

run-app:
	cd cmd/ordersystem && go run main.go wire_gen.go

run-evans:
	evans -r repl

gen-protoc:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

gen-wire:
	cd cmd/ordersystem && wire

gen-graphql:
	go run github.com/99designs/gqlgen generate