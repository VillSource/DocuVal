package main

import "github.com/villsource/docuval/internal/docuval"

func main() {
	docuvalServer := docuval.New()
	docuvalServer.Start()
}
