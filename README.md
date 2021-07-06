# OpenTracing Tutorial - Go

## Installing

This tutorial was inspired by [opentracing-tutorial-go](https://github.com/yurishkuro/opentracing-tutorial/tree/master/go) and add some more lessons.

```
git clone https://github.com/gogoods/opentracing-tutorial.git
cd $GOPATH/src/github.com/gogoods/opentracing-tutorial
go mod download
```

## Lessons

* [Lesson 01 - Hello World](./lesson01)
  * Instantiate a Tracer
  * Create a simple trace
  * Annotate the trace
* [Lesson 02 - Context and Tracing Functions](./lesson02)
  * Trace individual functions
  * Combine multiple spans into a single trace
  * Propagate the in-process context
* [Lesson 03 - Tracing in HTTP Requests](./lesson03)
  * Trace a transaction across more than one microservice
  * Pass the context between processes using `Inject` and `Extract`
  * Apply OpenTracing-recommended tags
* [Lesson 04 - Baggage](./lesson04)
  * Understand distributed context propagation
  * Use baggage to pass data through the call graph
* [Lesson 05 -Tracing in GRPC Requests](./lesson05)
  * 在GRPC协议中的调用跟踪使用方法
  * GRPC中客户端和服务端一元拦截器的实现
* [Lesson 06 - Other features](./lesson06)
  * Debug in sampling  
  
## References
- https://github.com/yurishkuro/opentracing-tutorial/blob/master/go/README.md
- https://blog.csdn.net/liyunlong41/article/details/87932953
- https://blog.csdn.net/liyunlong41/article/details/88043604
- https://github.com/moxiaomomo/grpc-jaeger
- https://github.com/jaegertracing/jaeger-client-go

