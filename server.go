package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aweimeow/grpc-gateway-example/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

type Server struct {
	protos.UnimplementedAdminServiceServer
}

var (
	count uint32 = 0
	data map[uint32]*Employee

	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50050", "gRPC server endpoint")
)

func init() {
	data = make(map[uint32]*Employee)
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	protos.RegisterAdminServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}

func StartHttpReverseProxyServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := protos.RegisterAdminServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		fmt.Println(err)
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}

func (s *Server) NewEmployee(ctx context.Context, in *protos.EmployeeCreateRequest) (*protos.EmployeeCreateResponse, error) {
	var isSuccess bool = true
	var message string

	if in.Name == "" || in.Age == 0 {
		isSuccess = false
		message = "Employee data wasn't given"
		return &protos.EmployeeCreateResponse{IsSuccess: isSuccess, Message: message}, nil
	}

	newEmployee := &Employee{
		name: in.Name,
		gender: Gender(in.Gender),
		age: in.Age,
	}

	data[count] = newEmployee
	count = count + 1

	fmt.Printf("Employee data: %v", data)

	return &protos.EmployeeCreateResponse{IsSuccess: true, Message: fmt.Sprintf("Employee craeted: %s", newEmployee)}, nil
}