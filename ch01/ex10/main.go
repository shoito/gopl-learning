// 大量のデータを生成するウェブサイトを見つけなさい。
// 報告される時間が大きく変化するかを調べるために fetchall を 2 回続けて実行して、
// キャッシュされているかどうかを調査しなさい。毎回同じ内容を得ているでしょうか。
// fetchall を修正して、その出力をファイルへ保存するようにして調べられるようにしなさい。

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const filePath = "./fetch.out"

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // ゴルーチンを開始
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer f.Close()

	for range os.Args[1:] {
		f.WriteString(<-ch)
		f.WriteString("\n")
	}
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // ch チャネルへ送信
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
