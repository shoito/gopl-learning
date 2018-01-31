package bench

import (
	"strings"
)

var arr = []string{"abc", "efg"}

func ForLoop() {
	var s, sep string
	for i := 0; i < len(arr); i++ {
		s += sep + arr[i]
		sep = " "
	}
}

func StringsJoin() {
	strings.Join(arr[0:], " ")
}
