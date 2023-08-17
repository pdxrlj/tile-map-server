package opentracing

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type Tracer struct {
	Tracer   io.Closer
	error    []error
	rootSpan opentracing.Span
	spanCtx  context.Context
}

func Init(service string, collectionEndpoint string) *Tracer {
	cfg := config.Configuration{
		// 将采样频率设置为1，每一个span都记录，方便查看测试结果
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			// 将span发往jaeger-collector的服务地址 http://localhost:6831
			CollectorEndpoint: collectionEndpoint,
		},
	}
	t := &Tracer{}
	closer, err := cfg.InitGlobalTracer(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		t.error = append(t.error, err)
		return t
	}
	t.Tracer = closer
	return t
}

func (t *Tracer) Start(ctx context.Context) *Tracer {
	trace := opentracing.GlobalTracer()
	span := trace.StartSpan("tile-map")
	t.rootSpan = span
	spanCtx := opentracing.ContextWithSpan(ctx, span)
	t.spanCtx = spanCtx
	return t
}

func (t *Tracer) GetSpanCtx() context.Context {
	return t.spanCtx
}

func (t *Tracer) CloseRootSpan() {
	t.rootSpan.Finish()
}

func (t *Tracer) Close() {
	_ = t.Tracer.Close()
}
