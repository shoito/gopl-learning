package main

import (
	"testing"
)

var input uint64 = 0x1234567890ABCDEF
var output int

func BenchmarkPopCount(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		c = PopCount(input)
	}
	output = c
}

func BenchmarkPopCountByShift(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		c = PopCountByShift(input)
	}
	output = c
}
