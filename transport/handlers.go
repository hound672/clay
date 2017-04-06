package transport

import (
	"net/http"

	"google.golang.org/grpc"
)

// Service is a registerable collection of endpoints.
// These functions should be autogenerated by LZD Protobuf codegenerator.
type Service interface {
	GetDescription() ServiceDesc
}

// ServiceDesc is a description of an endpoints' collection.
// These functions should be autogenerated by LZD Protobuf codegenerator.
type ServiceDesc interface {
	RegisterGRPC(*grpc.Server)
	RegisterHTTP(Router)
	SwaggerDef() []byte
}

// Router routes HTTP requests around.
type Router interface {
	Handle(pattern string, h http.Handler)
	HandleFunc(pattern string, h http.HandlerFunc)
}
