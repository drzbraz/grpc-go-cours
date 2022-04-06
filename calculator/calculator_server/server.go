package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/drzbraz/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstNumber := req.GetCalculator().GetFirstNumber()
	secondNumber := req.GetCalculator().GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func main(){
    fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err!=nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(lis); err!=nil {
		log.Fatalf("Fail to serve: %v", err)
	}
}