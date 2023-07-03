package lttng

/*
#cgo LDFLAGS: -ldl -llttng-ust

#define TRACEPOINT_DEFINE
#include "k8s-tp.h"

void traceEvent(char* o_name, unsigned int name_length, char* o_ctx, unsigned int ctx_length) {
	tracepoint(k8s_ust, event, o_name, name_length, o_ctx, ctx_length);
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func ReportEvent(operationName, context string) {
	operationNameBytes := []byte(operationName)
	contextBytes := []byte(context)
	C.traceEvent(
		(*C.char)(unsafe.Pointer(&operationNameBytes[0])),
		C.uint(len(operationNameBytes)),
		(*C.char)(unsafe.Pointer(&contextBytes[0])),
		C.uint(len(contextBytes)),
	)
	runtime.KeepAlive(operationNameBytes)
	runtime.KeepAlive(contextBytes)
}
