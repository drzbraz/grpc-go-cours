package main

import (
	"context"
	"fmt"
	"log"

	"github.com/drzbraz/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main(){

	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil{
		log.Fatal("could not connect: %v",err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)


	
}

func doUnary(c calculatorpb.CalculatorServiceClient){ 
	fmt.Println("Starting to do a Unary RPC....")

	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{FirstNumber: 2, SecondNumber: 2},
	}

	res,err := c.Calculator(context.Background(),req)
	if err != nil{
		log.Fatal("could not %v",err)
	}

	log.Println("Response from Greet Server:", res.Result)
}