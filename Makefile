# Variables
GO=go

# Targets
default: run

run:
	$(GO) run ./docuval/cmd/server.go