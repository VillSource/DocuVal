default: build

build: env-spawn
	@echo "Using go version $(shell go version | cut -d' ' -f3)"
	@go build cmd/docuval/docuval.go

serve: env-spawn
	@go run cmd/docuval/docuval.go

env-spawn:
	@if [ ! -f .env ]; then \
		echo "Creating .env from example..."; \
		cp .env.example .env; \
		echo "Please edit .env and add the correct values."; \
	fi
