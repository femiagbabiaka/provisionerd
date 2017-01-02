package main

import (
	"encoding/json"
	"net/http"
	
	"golang.org/x/net/context"
	"github.com/go-kit/kit/endpoint"
)

func makeCreateMailerEndpoint(svc Provisionerd) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createMailerRequest)
		v, err := svc.AddVirtualMailer(req.VM)

		if err != nil {
			return createMailerResponse{v, err.Error()}, nil
		}
		
		return createMailerResponse{v, ""}, nil
	}
}

func decodeCreateMailerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createMailerRequest
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type createMailerRequest struct {
	VM VirtualMailer
}

type createMailerResponse struct {
	VM VirtualMailer `json:"data,omitempty"`
	Err string `json:"err,omitempty"`
}
