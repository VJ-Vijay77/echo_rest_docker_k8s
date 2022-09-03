PROJECT_BINARY=binary

up: build
	@echo "starting postgres container"
	docker compose down
	docker compose up -d
	@echo "Started"
	

down:
	@echo "stopping container"
	docker compose down
	

build:
	@echo "Building project Binary"
	env GOOS=linux CGO_ENABLED=0 go build -o ${PROJECT_BINARY} .