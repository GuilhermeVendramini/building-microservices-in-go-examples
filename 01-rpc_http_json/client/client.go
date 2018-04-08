package client

import (
	"log"
	"net/rpc"

	"github.com/GuilhermeVendramini/microservices-in-go/01-rpc_http_json/contract"
)

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
