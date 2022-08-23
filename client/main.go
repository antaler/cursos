package main

import (
	"context"
	"io"
	"log"
	"time"

	"gitlab.com/antaler/cursos/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}

	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)
	//	DoUnary(c)
	//	DoClientStreaming(c)
	//	DoServerStreaming(c)
	DoBidirectionalStreaming(c)

}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response from Server: %v", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q8t1",
			Answer:   "Azul",
			Question: "Color Asociado  a Golang",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "Google",
			Question: "Empresa que desarrollo el lenguaje Golang",
			TestId:   "t1",
		},
		{
			Id:       "q10t1",
			Answer:   "Backend",
			Question: "Especialidad de Golang",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestions(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for _, question := range questions {
		log.Println("Sending Question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server response: %v", msg)

}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Print("END")
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println(msg)

	}

}

func DoBidirectionalStreaming(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequest{
		Answer: "casa",
	}

	numberOfQuestions := 8

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
				break
			}
			log.Printf("Respuesta %v", res)
		}
		log.Print("XXXX")
		close(waitChannel)
	}()

	<-waitChannel
}
