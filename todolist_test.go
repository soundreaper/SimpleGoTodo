package main

import (
	"testing"
)

func BenchmarkCreate10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		continue
	}
}
