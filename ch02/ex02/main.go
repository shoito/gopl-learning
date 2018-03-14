// コマンドライン引数、もしくはコマンドライン引数が指定されなかった場合には
// 標準入力から数値を読み込む、cfに似た汎用単位変換プログラムを書きなさい。
// 各数値は、温度なら華氏と摂氏で、長さならフィートとメートルで、
// 重さならポンドとキログラムでといった具合に各種単位へ変換しなさい。

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shoito/gopl-learning/ch02/ex02/conv"
)

func main() {
	// temp, length, weight
	var t, v string
	if len(os.Args) == 3 {
		t, v = os.Args[1], os.Args[2]
	} else {
		if _, err := fmt.Scanln(&t, &v); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	num, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	convert(t, num)
}

func convert(t string, v float64) {
	switch t {
	case "temp":
		f := conv.Fahrenheit(v)
		c := conv.Celsius(v)
		fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
	case "length":
		f := conv.Feet(v)
		m := conv.Meter(v)
		fmt.Printf("%s = %s, %s = %s\n", f, conv.FToM(f), m, conv.MToF(m))
	case "weight":
		p := conv.Pound(v)
		k := conv.Kilogram(v)
		fmt.Printf("%s = %s, %s = %s\n", p, conv.PToK(p), k, conv.KToP(k))
	default:
		fmt.Fprintln(os.Stderr, "error!")
	}
}
