package vault

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

//using NewServeMux to build a http.Handler interface with routing on hash and validate
//NewServer is how we get Go Kit to give us an HTTP handler for each endpoint
func newHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/hash", httptransport.NewServer(
		ctx,
		endpoints.HashEndpoint,
		decodeHashRequest,
		encodeResponse,
	))
	m.Handle("/validate", httptransport.NewServer(
		ctx,
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	))
	return m
}
