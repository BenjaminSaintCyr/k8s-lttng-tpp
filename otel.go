package lttng

import (
	"context"
	"time"

	"sync/atomic"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type LttngTracerProvider struct{}

var _ oteltrace.TracerProvider = LttngTracerProvider{}

func (p LttngTracerProvider) Tracer(instrumentationName string, _ ...oteltrace.TracerOption) oteltrace.Tracer {
	return LttngTracer{
		instrumentationName: instrumentationName,
	}
}

type LttngTracer struct {
	instrumentationName string
	currentID           uint64
}

var _ oteltrace.Tracer = LttngTracer{}

func (t LttngTracer) Start(ctx context.Context, name string, _ ...oteltrace.SpanOption) (context.Context, oteltrace.Span) {
	start := time.Now()
	span := LttngSpan{
		traceID: t.currentID,
		start:   start,
	}
	ReportStartSpan(t.currentID, t.currentID, name, start)
	atomic.AddUint64(&t.currentID, 1)
	return oteltrace.ContextWithSpan(ctx, span), span
}

type LttngSpan struct {
	traceID uint64
	start   time.Time
}

var _ oteltrace.Span = LttngSpan{}

func (LttngSpan) SpanContext() oteltrace.SpanContext { return oteltrace.SpanContext{} }

func (LttngSpan) IsRecording() bool { return true }

func (LttngSpan) SetStatus(code codes.Code, msg string) {}

func (LttngSpan) SetError(bool) {}

func (LttngSpan) SetAttributes(...attribute.KeyValue) {}

func (s LttngSpan) End(...oteltrace.SpanOption) {
	ReportEndSpan(s.traceID, time.Since(s.start))
}

func (LttngSpan) RecordError(error, ...oteltrace.EventOption) {}

func (LttngSpan) Tracer() oteltrace.Tracer { return LttngTracer{} }

func (LttngSpan) AddEvent(string, ...oteltrace.EventOption) {}

func (LttngSpan) SetName(string) {}
