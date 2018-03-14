// 引数をビットシフトしながら最下位ビットの検査を64回繰り返すことで
// ビット数を数えるPopCountのバージョンを作成しなさい。
// テーブル参照を行うバージョンと性能を比較しなさい。

package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByShift(x uint64) int {
	n := 0
	for i := 0; i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x = x >> 1
	}
	return n
}
