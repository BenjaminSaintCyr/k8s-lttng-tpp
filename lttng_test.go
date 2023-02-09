package lttng

import "testing"

func BenchmarkReportStartSpan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportStartSpan(1, 0, "test", "testing")
	}
}

func BenchmarkReportEndSpan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportEndSpan(1, "testing")
	}
}
