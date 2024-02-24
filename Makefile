.PHONY: dev build

dev:
	@echo "Starting development environment..."
	@{ \
	trap 'trap - SIGTERM && kill -- -$$' SIGINT SIGTERM EXIT; \
	cd client && yarn dev & \
	cd server && go run -tags dev . & \
	wait; \
	}

build:
	@echo "Building client..."
	@cd client && yarn build
	@echo "Building Go server for production..."
	@cd server && go build -o ../myapp -tags !dev .
	@echo "Build complete. Execute ./myapp to run the server."
