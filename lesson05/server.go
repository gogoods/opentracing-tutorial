package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	pb "samples/jaeger/lesson05/proto"
	"samples/jaeger/lib/tracing"
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpclib "samples/jaeger/lib/grpc"
)

const (
	//gRPC服务地址
	Address = "0.0.0.0:5050"
)

type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {



	resp := new(pb.HelloReply)
	resp.Greet = "Bonjour"

	return resp, nil
}

var HelloServer = helloService{}

func main() {

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Panicf("failed to listen:%v", err)
	}

	var servOpts []grpc.ServerOption
	tracer, closer := tracing.NewJaegerTracer("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	servOpts = append(servOpts, grpclib.ServerOption(tracer))

	//实现gRPC Server
	s := grpc.NewServer(
		servOpts...,
		)

	fmt.Println(s.GetServiceInfo())

	defer s.GracefulStop()
	pb.RegisterGreeterServer(s, HelloServer)

	reflection.Register(s)
	fmt.Println("Listen on " + Address)

	s.Serve(listen)
}

