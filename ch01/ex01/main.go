// gopl.io/ch1/echo1 を修正して、そのプログラムを起動したコマンド名であるos.Args[0]も表示するようにしなさい。

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
