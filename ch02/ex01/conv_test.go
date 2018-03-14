// 絶対温度(Kelvin scale)で温度を処理するためにtempconvに型、定数、関数を追加しなさい。
// 0Kは-273.15℃であり、1Kの差と1℃の差は同じ大きさです。

package main

import (
	"math"
	"testing"
)

func TestCToK(t *testing.T) {
	if k := CToK(-273.15); !equals(float64(k), 0) {
		t.Errorf("-273.15°C must be 0K, but %f", k)
	}
	if k := CToK(0); !equals(float64(k), 273.15) {
		t.Errorf("0°C must be 273.15K, but %f", k)
	}
}

func TestKToC(t *testing.T) {
	if c := KToC(273.15); !equals(float64(c), 0) {
		t.Errorf("273.15K must be 0°C, but %f", c)
	}
	if c := KToC(0); !equals(float64(c), -273.15) {
		t.Errorf("0K must be -273.15°C, but %f", c)
	}
}

func TestFToK(t *testing.T) {
	if k := FToK(212); !equals(float64(k), 373.15) {
		t.Errorf("212°F must be 373.15K, but %f", k)
	}
	if k := FToK(32); !equals(float64(k), 273.15) {
		t.Errorf("32°F must be 273.15K, but %f", k)
	}
}

func TestKToF(t *testing.T) {
	if f := KToF(273.15); !equals(float64(f), 32) {
		t.Errorf("273.15K must be 32°F, but %f", f)
	}
	if f := KToF(0); !equals(float64(f), -459.67) {
		t.Errorf("0K must be -459.67°F, but %f", f)
	}
}

func equals(a, b float64) bool {
	return math.Abs(float64(a)-float64(b)) < 0.1
}
