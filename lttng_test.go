package lttng

import "testing"

func BenchmarkReportEvent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportEvent("test", "testing")
	}
}
