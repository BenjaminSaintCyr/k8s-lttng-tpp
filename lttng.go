package lttng

/*
#cgo LDFLAGS: -ldl -llttng-ust

#define TRACEPOINT_DEFINE
#include "k8s-tp.h"

void traceStartSpan(uint64_t s_id, uint64_t s_p_id, char* o_name, char* o_ctx) {
	tracepoint(k8s_ust, start_span, s_id, s_p_id, o_name, o_ctx);
}

void traceEndSpan(uint64_t s_id, char* o_ctx) {
	tracepoint(k8s_ust, end_span, s_id, o_ctx);
}
*/
import "C"

func ReportStartSpan(spanID uint64, parentID uint64, operationName string, context string) {
	C.traceStartSpan(
		C.uint64_t(spanID),
		C.uint64_t(parentID),
		C.CString(operationName),
		C.CString(context),
	)
}

func ReportEndSpan(spanID uint64, context string) {
	C.traceEndSpan(
		C.uint64_t(spanID),
		C.CString(context),
	)
}
