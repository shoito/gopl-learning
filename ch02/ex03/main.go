// 単一の式の代わりにループを使うようにPopCountを書き直しなさい。
// 2つのバージョンの性能を比較しなさい。
// (11.4節で異なる実装の性能を体系的に比較する方法を説明しています。)

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

func PopCountByLoop(x uint64) int {
	var n int
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(uint(i)*8))])
	}
	return n
}
