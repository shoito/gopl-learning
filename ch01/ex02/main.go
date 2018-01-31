// gopl.io/ch1/echo1 を修正して、個々の引数のインデックスと値の組を1行ごとに表示しなさい。

package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d:%s\n", i, os.Args[i])
	}
}
