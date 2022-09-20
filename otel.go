package lttng

import (
	"context"
	"time"

	"sync/atomic"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type lttngTracerProvider struct{}

var _ oteltrace.TracerProvider = lttngTracerProvider{}

func (p lttngTracerProvider) Tracer(instrumentationName string, _ ...oteltrace.TracerOption) oteltrace.Tracer {
	return lttngTracer{
		instrumentationName: instrumentationName,
	}
}

type lttngTracer struct {
	instrumentationName string
	currentID           uint64
}

var _ oteltrace.Tracer = lttngTracer{}

func (t lttngTracer) Start(ctx context.Context, name string, _ ...oteltrace.SpanOption) (context.Context, oteltrace.Span) {
	start := time.Now()
	span := lttngSpan{
		traceID: t.currentID,
		start:   start,
	}
	ReportStartSpan(t.currentID, t.currentID, t.currentID, t.currentID, name, start)
	atomic.AddUint64(&t.currentID, 1)
	return oteltrace.ContextWithSpan(ctx, span), span
}

type lttngSpan struct {
	traceID uint64
	start   time.Time
}

var _ oteltrace.Span = lttngSpan{}

func (lttngSpan) SpanContext() oteltrace.SpanContext { return oteltrace.SpanContext{} }

func (lttngSpan) IsRecording() bool { return true }

func (lttngSpan) SetStatus(code codes.Code, msg string) {}

func (lttngSpan) SetError(bool) {}

func (lttngSpan) SetAttributes(...attribute.KeyValue) {}

func (s lttngSpan) End(...oteltrace.SpanOption) {
	ReportEndSpan(s.traceID, s.traceID, s.traceID, time.Since(s.start))
}

func (lttngSpan) RecordError(error, ...oteltrace.EventOption) {}

func (lttngSpan) Tracer() oteltrace.Tracer { return lttngTracer{} }

func (lttngSpan) AddEvent(string, ...oteltrace.EventOption) {}

func (lttngSpan) SetName(string) {}
