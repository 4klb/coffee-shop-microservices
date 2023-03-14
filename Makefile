.PHONY: gen-dish
gen-dish:
	protoc -I proto/dish --go_out=proto dish.proto

.PHONY: gen-dishService
gen-dishService:
	protoc -I proto --go_out=proto --go-grpc_out=proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative dish/dish.proto
	