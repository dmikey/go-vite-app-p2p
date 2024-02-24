.PHONY: dev build proto build-app

dev:
	@echo "Starting development environment..."
	@{ \
	trap 'trap - SIGTERM && kill -- -$$' SIGINT SIGTERM EXIT; \
	cd client && yarn dev > /dev/null & \
	cd server && go run -tags dev . & \
	wait; \
	}

build:
	@echo "Building client..."
	@cd client && yarn build
	@echo "Building Go server for production..."
	@cd server && go build -o ../myapp -tags !dev,!app .
	@echo "Build complete. Execute ./myapp to run the server."

proto-gen:
	@echo "Generating Go code from .proto files..."
	@protoc --go_out=paths=source_relative:./server --go_opt=paths=source_relative ./proto/*.proto
	@protoc --proto_path=./proto --js_out=import_style=es6,binary:./client/src ./proto/*.proto
	@echo "Proto compilation complete."