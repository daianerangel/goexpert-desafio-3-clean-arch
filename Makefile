docker-up:
	docker-compose up -d

run:
	cd cmd/ordersystem && go run main.go wire_gen.go

gen-protoc:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

gen-wire:
	cd cmd/ordersystem && wire

run-evans:
	evans -r repl

build-app:
	cd cmd/ordersystem && go build -o app -v