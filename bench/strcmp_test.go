package bench

import "testing"

var result bool

func BenchmarkCompareLen(b *testing.B) {

	for n := 0; n < b.N; n++ {
		CompareLen("")
	}
}

func BenchmarkCompareStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CompareStr("")
	}
}

func BenchmarkCompareLenLongStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CompareLen("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	}
}

func BenchmarkCompareStrLongStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CompareStr("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	}
}
