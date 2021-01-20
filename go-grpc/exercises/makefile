compile:
	protoc --go_out=plugins=grpc:client/ --go_opt=paths=source_relative \
	--go_out=plugins=grpc:server/ --go_opt=paths=source_relative proto/order/order.proto \