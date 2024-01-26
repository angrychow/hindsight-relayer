package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"net"

	traceHelper "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/grpc"
)

type RelayTraceServiceServer struct {
	traceHelper.UnimplementedTraceServiceServer
}

func (RelayTraceServiceServer) Export(
	ctx context.Context,
	request *traceHelper.ExportTraceServiceRequest) (
	*traceHelper.ExportTraceServiceResponse, error) {
	fmt.Println(hex.EncodeToString(request.ResourceSpans[0].ScopeSpans[0].Spans[0].SpanId))
	resp := new(traceHelper.ExportTraceServiceResponse)
	return resp, nil
}

func main() {
	// lis, err := net.Listen("tcp", ":50051")
	serv := grpc.NewServer()
	traceHelper.RegisterTraceServiceServer(serv, RelayTraceServiceServer{})
	listenAddr := "0.0.0.0:4317"
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Server listening on %s", listenAddr)
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
