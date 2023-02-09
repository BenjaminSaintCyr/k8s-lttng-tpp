package lttng

import (
	"sync/atomic"
)

type LttngCtx struct {
	Id uint64
}

var (
	IDcounter uint64 = 1
)

func ReportStart(operationName, context string) LttngCtx {
	atomic.AddUint64(&IDcounter, 1)
	id := IDcounter
	var parent uint64 = 0

	ReportStartSpan(id, parent, operationName, context)
	return LttngCtx{
		Id: id,
	}
}

func (ctx *LttngCtx) End(context string) {
	ReportEndSpan(ctx.Id, context)
}

func (ctx *LttngCtx) ReportChild(operationName, context string) LttngCtx {
	atomic.AddUint64(&IDcounter, 1)
	id := IDcounter

	ReportStartSpan(id, ctx.Id, operationName, context)
	return LttngCtx{
		Id: id,
	}
}
