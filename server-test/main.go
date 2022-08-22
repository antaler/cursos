package main

import (
	"log"
	"net"

	"gitlab.com/antaler/cursos/database"
	"gitlab.com/antaler/cursos/server"
	"gitlab.com/antaler/cursos/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatal(err)
	}
	url := "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"
	repo, err := database.NewPostgresRepository(url)

	server := server.NewTestServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
