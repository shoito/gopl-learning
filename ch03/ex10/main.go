// 文字列の連結の代わりにbytes.Bufferを使って、再帰呼出しを行わないcommaを作成しなさい

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("comma: %s\n", comma(os.Args[i]))
		fmt.Printf("buff:  %s\n", buff(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func buff(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	init := n % 3
	if init == 0 {
		init = 3
	}

	var b bytes.Buffer
	b.WriteString(s[:init])
	for i := init; (i + 3) <= n; i += 3 {
		b.WriteString(",")
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
