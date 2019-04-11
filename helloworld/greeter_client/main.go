/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
//	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/sanshirookazaki/grpc-sample/helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func printGreeting(client pb.GreeterClient, rep *pb.HelloRequest) {
	log.Printf("greeting::: %v", rep)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.SayHello(ctx, rep)
	if err != nil {
		log.Fatalf("%v.greeting(_) = _, %v", client, err)
	}
	for {
		greeting, err := stream.Recv()
		if err == io.EOF {
				break
		}
		if err != nil {
			log.Fatalf("%v.greeting(_) = _, %v", client, err)
		}
		log.Println(greeting)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	var age int64
	age = 25
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Age: age})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.Message)
	//fmt.Println(int(r.Age))
	printGreeting(c, &pb.HelloRequest{Name: name, Age: age})
	printGreeting(c, &pb.HelloRequest{Name: "test", Age: 15})
}
