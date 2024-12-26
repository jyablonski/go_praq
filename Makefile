.PHONY test:
test:
	@docker compose -f docker/docker-compose-test.yml up -d

.PHONY: up
up:
	@docker compose -f docker/docker-compose-postgres.yml up -d postgres

.PHONY: down
down:
	@docker compose -f docker/docker-compose-postgres.yml down

# grpc cokmmands
run-orders:
	@go run services/orders/*.go

run-kitchen:
	@go run services/kitchen/*.go

.PHONY: gen
gen:
	@protoc \
		--proto_path=grpc_project/protobuf "grpc_project/protobuf/orders.proto" \
		--go_out=grpc_project/services/common/genproto/orders --go_opt=paths=source_relative \
  		--go-grpc_out=grpc_project/services/common/genproto/orders --go-grpc_opt=paths=source_relative