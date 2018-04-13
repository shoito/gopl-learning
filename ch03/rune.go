package main

import "fmt"

func main() {
	runes := []rune("𠮷野家")
	for _, r := range runes {
		fmt.Printf("%U %s\n", r, string(r))
	}
}