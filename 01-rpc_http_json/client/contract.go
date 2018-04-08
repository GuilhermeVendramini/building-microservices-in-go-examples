package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/GuilhermeVendramini/microservices-in-go/01-rpc_http_json/contract"
)

// PerformRequest ...
func PerformRequest() contract.HelloWorldResponse {
	r, _ := http.Post(
		"http://localhost:1234",
		"application/json",
		bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
	)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var response contract.HelloWorldResponse
	decoder.Decode(&response)

	return response
}
