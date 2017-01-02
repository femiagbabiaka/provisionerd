package main

import (
	"golang.org/x/net/context"
	"os"
	"net/http"
	
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
)
	
func main() {
	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	
	var svc Provisionerd
	svc = provisionerd{}
	
	addVirtualMailerHandler := httptransport.NewServer(
		ctx,
		makeCreateMailerEndpoint(svc),
		decodeCreateMailerRequest,
		encodeResponse,
	)
	
	r := mux.NewRouter()
	r.Methods("POST").Path("/virtualmailer").Handler(addVirtualMailerHandler)
	
	http.Handle("/", r)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
