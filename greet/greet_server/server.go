package main
import (
	"fmt"
	"log"
	"net"

	"github.com/drzbraz/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func main(){
    fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err!=nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err!=nil {
		log.Fatalf("Fail to serve: %v", err)
	}
}