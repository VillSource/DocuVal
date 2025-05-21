package server

import "fmt"

type DocuvalServer struct {}

func New() *DocuvalServer {
	return &DocuvalServer{}
}

func (s *DocuvalServer) Start() {
	fmt.Println("Docuval server is running...")
}
