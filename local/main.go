package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"samples/jaeger/lib/tracing"
	"time"
)

func func1(req string, ctx context.Context) (reply string){
	//1.创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "func1")
	defer func() {
		//4.接口调用完，在tag中设置request和reply

		span.SetTag("request", req)
		span.SetTag("response", reply)
		span.Finish()
	}()

	println(req)
	//2.模拟处理耗时
	time.Sleep(time.Second/2)
	//3.返回reply
	reply = "func1 result"
	return
}

func func2(req string, ctx context.Context) (reply string){
	span, newCtx := opentracing.StartSpanFromContext(ctx, "func2")
	defer func() {
		span.SetTag("request", req)
		span.SetTag("response", reply)
		span.Finish()
	}()

	println(req)
	time.Sleep(time.Second/2)


	// 子调用
	//r3 := func3("Hello func3", opentracing.ContextWithSpan(ctx, span))
	r3 := func3("Hello func3", newCtx)
	fmt.Println(r3)

	reply = "func2 result" + " "+ r3
	return
}

func func3(req string, ctx context.Context) (reply string){
	span, _ := opentracing.StartSpanFromContext(ctx, "func3")
	defer func() {
		span.SetTag("request", req)
		span.SetTag("response", reply)
		span.Finish()
	}()

	println(req)
	time.Sleep(time.Second/2)
	reply = "func3 result"
	return
}

func main() {
	tracer, closer := tracing.NewJaegerTracer("jaeger-demo")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)//StartspanFromContext创建新span时会用到

	// 创建根trace
	root := tracer.StartSpan("root")

	ctx := opentracing.ContextWithSpan(context.Background(), root)

	r1 := func1("Hello func1", ctx)
	fmt.Println(r1)

	r2 := func2("Hello func2", ctx)
	fmt.Println(r2)

	root.Finish()
}
