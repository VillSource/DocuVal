package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/villsource/docuval-identity/pkg/identity"
)

type DocuvalServer struct {
    fiberApp *fiber.App
}

func New() *DocuvalServer {
	return &DocuvalServer{
        fiberApp: fiber.New(),
    }
}

func (s *DocuvalServer) Start() {
	fmt.Println("Docuval server is running...")

    s.configIdentity()

    fmt.Println("http://localhost:3000")
    s.fiberApp.Listen(":3000")
}


func (s *DocuvalServer) configIdentity() {
    s.fiberApp.Use(identity.NewFiberMiddleware())
}