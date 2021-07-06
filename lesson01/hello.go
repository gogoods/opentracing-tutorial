package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"os"
	"samples/jaeger/lib/tracing"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	helloTo := os.Args[1]
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)

	//exercise2(helloStr)
	//exercise2(helloTo, helloStr)
	//exercise3(helloTo, helloStr)
	exercise4(helloTo, helloStr)
}

// Local tracer
func exercise1(v string){
	tracer := opentracing.GlobalTracer()

	span := tracer.StartSpan("say-hello")

	println(v)
	span.Finish()
}

// Remote tracer
func exercise2(v string){
	tracer, closer := tracing.NewJaegerTracer("hello-world")
	defer closer.Close()

	span := tracer.StartSpan("say-hello")


	println(v)
	span.Finish()
}

// Remote tracer and annotate the Trace with Tags
func exercise3(v1, v2 string){
	tracer, closer := tracing.NewJaegerTracer("hello-world")
	defer closer.Close()

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", v1)
	println(v2)
	span.Finish()
}

// Remote tracer and annotate the Trace with Tags and Logs
func exercise4(v1, v2 string){
	tracer, closer := tracing.NewJaegerTracer("hello-world")
	defer closer.Close()

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", v1)

	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", v2),
	)

	println(v2)
	span.LogKV("event", "println", "value", v2)
	span.Finish()
}