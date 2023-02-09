package lttng

import "testing"

func BenchmarkReportStart(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportStart("test", "testing")
	}
}

func BenchmarkEnd(b *testing.B) {
	ctx := LttngCtx{
		Id: uint64(0),
	}
	for n := 0; n < b.N; n++ {
		ctx.End("end test")
	}
}

func BenchmarkReportChild(b *testing.B) {
	ctx := LttngCtx{
		Id: uint64(0),
	}
	for n := 0; n < b.N; n++ {
		ctx.ReportChild("test", "testing")
	}
}
