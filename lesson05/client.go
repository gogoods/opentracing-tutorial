package main

import (
	"flag"
	"fmt"
	"github.com/opentracing/opentracing-go"
	pb "samples/jaeger/lesson05/proto"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"samples/jaeger/lib/tracing"
	grpclib "samples/jaeger/lib/grpc"
)

func main() {
	var target = flag.String("target", "127.0.0.1:5050", "grpc server address")
	flag.Parse()

	log.Printf("-> target: %s", *target )

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	tracer, closer := tracing.NewJaegerTracer("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	dialOpts = append(dialOpts, grpclib.DialOption(tracer))

	conn, err := grpc.Dial(*target, dialOpts...)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	callGreeter(conn)
}



func callGreeter(conn *grpc.ClientConn) {
	c := pb.NewGreeterClient(conn)
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "client01"

	md := metadata.New(map[string]string{"uid": "99101691"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	var header, trailer metadata.MD
	rsp, err := c.SayHello(ctx, reqBody, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("header:", header)
	log.Println("trailer:", trailer)
	log.Println(rsp.Greet)
}


