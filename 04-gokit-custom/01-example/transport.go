package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

func makeGetEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		v, err := svc.Get(req.S)
		if err != nil {
			return getResponse{v, err.Error()}, nil
		}
		return getResponse{v, ""}, nil
	}
}

func makeDeleteEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequest)
		i, err := strconv.Atoi(req.S)
		if err != nil {
			return nil, err
		}
		v, err := svc.Delete(i)
		return deleteResponse{v}, err
	}
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getRequest struct {
	S string `json:"s"`
}

type getResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type deleteRequest struct {
	S string `json:"s"`
}

type deleteResponse struct {
	V bool `json:"v"`
}
