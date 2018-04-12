// できるだけコンパクトにKB、MB、...、YBまでのconst宣言を書きなさい

package main

import "fmt"

const (
	B  = 1024
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PG = TB * 1024
	EB = PG * 1024
	ZB = EB * 1024
	YB = ZB * 1024
)

func main() {
	fmt.Println(B, KB, MB)
}
