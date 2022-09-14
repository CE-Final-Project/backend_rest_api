
# ==============================================================================
# Go migrate postgresql https://github.com/golang-migrate/migrate

DB_NAME = account_db
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable
DB_USER = admin
DB_PASS = "2iNyADya#bUu9"

force_db:
	migrate -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations force 1

version_db:
	migrate -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations version

migrate_up:
	migrate -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations up 1

migrate_down:
	migrate -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations down 1


proto_kafka:
	@echo Generating kafka proto
	cd proto/kafka && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. kafka.proto

proto_account:
	@echo Generating account microservice proto
	cd account_service/proto/account && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. account.proto

proto_account_message:
	@echo Generating account messages microservice proto
	cd account_service/proto/account && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. account_messages.proto
