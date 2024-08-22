# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: run the application
.PHONY: run
run:
	@go run ./cmd/parser

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...' go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	#staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build: build the application
.PHONY: build
build:
	@echo 'Building cmd/parser...'
	go build -ldflags "-w -s" -o=./bin/parser ./cmd/parser
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o=./bin/linux_amd64/parser ./cmd/parser
	GOOS=linux GOARCH=arm64 go build -ldflags "-w -s" -o=./bin/linux_arm64/parser ./cmd/parser

## build: docker build local image
.PHONY: docker/build
docker/build:
	docker build --no-cache -t mvpotter/trstmtparser:v0.0.1-local .

## build: docker push image
.PHONY: docker/push
docker/push:
	docker buildx build --platform linux/amd64,linux/arm64 --push -t mvpotter/trstmtparser:v0.0.1 .
