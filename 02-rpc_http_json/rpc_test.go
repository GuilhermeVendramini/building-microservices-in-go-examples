package main

import (
	"testing"

	"github.com/GuilhermeVendramini/microservices-in-go/02-rpc_http_json/client"
	"github.com/GuilhermeVendramini/microservices-in-go/02-rpc_http_json/server"
)

func BenchmarkHelloWorldHandlerJSONRPC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest()

	}
}

func init() {
	// start the server
	go server.StartServer()
}
