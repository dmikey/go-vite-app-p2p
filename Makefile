.PHONY: dev build proto build-osx build-client

dev:
	@echo "Starting development environment..."
	@{ \
	trap 'trap - SIGTERM && kill -- -$$' SIGINT SIGTERM EXIT; \
	cd client && yarn dev > /dev/null & \
	cd node && go run -tags dev . & \
	wait; \
	}

build-client:
	@echo "Building client..."
	@cd client && yarn && yarn build

build-node:
	@echo "Building Go node for production..."
	@cd node && go build -o ../myapp -tags !dev,!app .
	@echo "Build complete. Execute ./myapp to run the node."

build:
	@echo "Building client..."
	@cd client && yarn && yarn build
	@echo "Building Go node for production..."
	@cd node && go build -o ../myapp -tags !dev,!app .
	@echo "Build complete. Execute ./myapp to run the node."

build-osx:
	@echo "Building client..."
	@cd client && yarn && yarn build
	@echo "Building OSX Shell..."
	@echo "Building Go node for production..."
	@cd node && go build -o ../myapp -tags !dev,!app .
	@cd ./containers/osx/SwiftFrame && xcodebuild -project SwiftFrame.xcodeproj -configuration Debug
	@echo "Build complete. Execute ./myapp to run the node."

proto-gen:
	@echo "Generating Go code from .proto files..."
	@protoc --go_out=paths=source_relative:./node --go_opt=paths=source_relative ./proto/*.proto
	@protoc --plugin=client/node_modules/ts-proto/protoc-gen-ts_proto --ts_proto_out=./client/src ./proto/*.proto --ts_proto_opt=esModuleInterop=true
	@echo "Proto compilation complete."