package lttng

import (
	"sync/atomic"
)

type Ctx struct {
	ID uint64
}

var (
	idCounter uint64 = 1
)

func ReportStart(operationName, context string) Ctx {
	id := atomic.AddUint64(&idCounter, 1)
	const parent uint64 = 0

	ReportStartSpan(id, parent, operationName, context)
	return Ctx{
		ID: id,
	}
}

func (ctx *Ctx) End(context string) {
	ReportEndSpan(ctx.ID, context)
}

func (ctx *Ctx) ReportChild(operationName, context string) Ctx {
	atomic.AddUint64(&idCounter, 1)
	id := idCounter

	ReportStartSpan(id, ctx.ID, operationName, context)
	return Ctx{
		ID: id,
	}
}
