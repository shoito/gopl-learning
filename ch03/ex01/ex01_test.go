package main

import (
	"math"
	"testing"
)

func TestF(t *testing.T) {
	x := f(float64(0), float64(0))
	if math.IsNaN(x) {
		t.Error("NaNですよ")
	}

	x = f(float64(10), float64(20))
	if math.IsNaN(x) {
		t.Error("NaNですよ")
	}
}
