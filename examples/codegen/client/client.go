package main

import (
	"context"
	"fmt"
	"time"

	"github.com/restatedev/sdk-go/ingress"
	helloworld "github.com/restatedev/sdk-go/examples/codegen/proto"
)

func main() {
	client := ingress.NewClient("http://127.0.0.1:8080")

	// Example with custom context (e.g., with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	greeter := helloworld.NewGreeterIngressClient(client)
	greeting, err := greeter.SayHello().Request(ctx, &helloworld.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting.Message)

	workflow := helloworld.NewWorkflowIngressClient(client, "my-workflow")
	invocation := workflow.Run().Send(context.Background(), &helloworld.RunRequest{})
	
	status, err := workflow.Status().Request(context.Background(), &helloworld.StatusRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("workflow running with invocation id", invocation.Id, "and status", status.Status)

	if _, err := workflow.Finish().Request(context.Background(), &helloworld.FinishRequest{}); err != nil {
		panic(err)
	}
}
