package main

import "testing"

func t(b *testing.B, f func(int, string) string) {
	str := randStr(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B) {
	t(b, plusConcat)
}

func BenchmarkFmtSprintf(b *testing.B) {
	t(b, fmtSprintf)
}

func BenchmarkStrBuilder(b *testing.B) {
	t(b, strBuilder)
}

func BenchmarkStrBuffer(b *testing.B) {
	t(b, strBuffer)
}

func BenchmarkStrBytes(b *testing.B) {
	t(b, strBytes)
}
