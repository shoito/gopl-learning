// 入力テキストファイル内のそれぞれの単語の出現頻度を報告する
// プログラムwordfreqを書きなさい。入力を行ではなく単語へ分割するために、
// 最初のScan呼び出しの前にinput.Split(bufio.ScanWords)を呼び出しなさい。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("word\tfreq\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
