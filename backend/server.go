package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync/atomic"

	// Ensure correct import paths
	counterv1 "crimtech_deliverable/proto/proto"
	counterv2 "crimtech_deliverable/proto/proto/protoconnect"

	"github.com/bufbuild/connect-go"
	"github.com/rs/cors"
)

// Atomic global counter
var globalCounter int64

// Implement the CounterServiceServer
type CounterServiceServer struct{}

// The Increment method that gets called by the frontend
func (s *CounterServiceServer) Increment(
	ctx context.Context,
	req *connect.Request[counterv1.IncrementRequest],
) (*connect.Response[counterv1.IncrementResponse], error) {

	// increment the counter
	newValue := atomic.AddInt64(&globalCounter, 1)

	// log each increment
	fmt.Println("Received increment request. New counter value:", newValue)

	// return the new counter value
	return connect.NewResponse(&counterv1.IncrementResponse{
		NewValue: newValue,
	}), nil
}

// Start the server 
func startServer() {
	// Create a new HTTP mux 
	mux := http.NewServeMux()

	// Register the ConnectRPC service handler
	path, handler := counterv2.NewCounterServiceHandler(&CounterServiceServer{})
	mux.Handle(path, handler)

	// Enable CORS 
	corsHandler := cors.Default().Handler(mux)

	// Listen for incoming requests 
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("✅ Server is running on http://localhost:8080")
	log.Fatal(http.Serve(listener, corsHandler))
}

func main() {
	startServer()
}
