package main

import (
	"log"
	"net"

	"gitlab.com/antaler/cursos/database"
	"gitlab.com/antaler/cursos/server"
	"gitlab.com/antaler/cursos/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatal(err)
	}
	url := "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"
	repo, err := database.NewPostgresRepository(url)

	server := server.NewStudentServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
