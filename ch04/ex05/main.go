// []stringスライス内で隣接している重複をスライス内で除去する関数を書きなさい。

package main

import "fmt"

func main() {
	str := []string{"a", "a", "a", "b", "b"}
	fmt.Printf("%v\n", unique(str))
}

func unique(str []string) []string {
	i := 0
	for _, s := range str {
		if str[i] == s {
			continue
		}
		i++
		str[i] = s
	}
	return str[:i+1]
}
