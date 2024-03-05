// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 	"google.golang.org/grpc"

// 	pb "./pb/data.pb.go" // Замените на фактический путь вашего пакета
// )

// func main() {
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterYourServiceServer(grpcServer, &YourServiceImplementation{}) // Замените на вашу реализацию gRPC-сервиса

// 	mux := runtime.NewServeMux()
// 	err := pb.RegisterYourServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8080", []grpc.DialOption{grpc.WithInsecure()})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	http.Handle("/", mux)

// 	log.Println("Server listening on :8080...")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
