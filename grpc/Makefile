pb-all:
	@make pb-user && make pb-product && make pb-cart && make pb-order && make pb-payment

pb-user:
	@cd user && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative error.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative handler.proto --go-grpc_opt=require_unimplemented_servers=false

pb-product:
	@cd product && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative product.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative error.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative handler.proto --go-grpc_opt=require_unimplemented_servers=false
	
pb-cart:
	@cd cart && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cart.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative error.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative handler.proto --go-grpc_opt=require_unimplemented_servers=false

pb-order:
	@cd order && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative order.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative error.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative handler.proto --go-grpc_opt=require_unimplemented_servers=false

pb-payment:
	@cd payment && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative payment.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative error.proto --go-grpc_opt=require_unimplemented_servers=false && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative handler.proto --go-grpc_opt=require_unimplemented_servers=false
