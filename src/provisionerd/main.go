package main

import (
	"net/http"
	"golang.org/x/net/context"
	"os"

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
	
	http.Handle("/addVirtualMailer", addVirtualMailerHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
