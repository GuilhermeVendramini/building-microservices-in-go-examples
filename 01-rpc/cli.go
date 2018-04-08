package main

import (
	"fmt"

	"github.com/GuilhermeVendramini/microservices-in-go/01-rpc/client"
	"github.com/GuilhermeVendramini/microservices-in-go/01-rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
