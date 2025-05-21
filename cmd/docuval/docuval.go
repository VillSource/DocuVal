package main

import (
	"fmt"

	"github.com/villsource/docuval/internal/server"
)

func main() {
	fmt.Println("Hello")
	docuvalServer := server.New()
	docuvalServer.Start()
}
