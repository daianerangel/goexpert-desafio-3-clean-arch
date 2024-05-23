docker-up:
	docker-compose up -d

run:
	cd cmd/ordersystem && go run main.go wire_gen.go