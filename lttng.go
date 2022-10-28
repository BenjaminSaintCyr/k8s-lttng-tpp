package lttng

/*
#cgo LDFLAGS: -ldl -llttng-ust

#define TRACEPOINT_DEFINE
#include "k8s-tp.h"

void traceStartSpan(uint64_t s_id, uint64_t s_p_id, char* o_name, int64_t s_time) {
	tracepoint(k8s_ust, start_span, s_id, s_p_id, o_name, s_time);
}

void traceEndSpan(uint64_t s_id, int64_t dur) {
	tracepoint(k8s_ust, end_span, s_id, dur);
}
*/
import "C"

import (
	"sync/atomic"
	"time"
)

var IDcounter uint64 = 1

func ReportStartSpan(spanID uint64, parentID uint64, operationName string, startTime time.Time) {
	C.traceStartSpan(
		C.uint64_t(spanID),
		C.uint64_t(parentID),
		C.CString(operationName),
		C.int64_t(startTime.UnixNano()),
	)
}

func ReportEndSpan(spanID uint64, duration time.Duration) {
	C.traceEndSpan(
		C.uint64_t(spanID),
		C.int64_t(duration.Nanoseconds()),
	)
}
