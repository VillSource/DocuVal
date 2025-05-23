default: build

build:
	@go build cmd/docuval/docuval.go

serve:
	@go run cmd/docuval/docuval.go
