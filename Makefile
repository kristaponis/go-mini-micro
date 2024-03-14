## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "Docker containers started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker-compose
up_build: build_mainApp build_authApp
	@echo "Stopping docker containers (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker containers..."
	docker-compose up --build -d
	@echo "Docker containers built and started!"

## down: stops docker-compose
down:
	@echo "Stopping docker-compose..."
	docker-compose down
	@echo "Done!"

## build_mainApp: builds the mainApp binary as a linux executable
build_mainApp:
	@echo "Building mainApp binary..."
	cd ./mainapp && env GOOS=linux CGO_ENABLED=0 go build -o mainApp
	@echo "Done!"

## build_authApp: builds the authApp binary as a linux executable
build_authApp:
	@echo "Building mainApp binary..."
	cd ./auth && env GOOS=linux CGO_ENABLED=0 go build -o authApp
	@echo "Done!"

## build_front: builds the frontend binary
build_front:
	@echo "Building front end binary..."
	cd ./frontend && env CGO_ENABLED=0 go build -o frontApp
	@echo "Done!"

## start: starts the frontend
start: build_front
	@echo "Starting frontend"
	cd ./frontend && ./frontApp

## stop: stops the frontend
stop:
	@echo "Stopping frontend..."
	@-pkill -SIGTERM -f "./frontApp"
	@echo "Stopped frontend!"