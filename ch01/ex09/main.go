// fetch を修正して、resp.Status に設定されている HTTP ステータスコードも表示するようにしなさい。

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const urlPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, urlPrefix) {
			url = urlPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Body:%s\nStatus:%s", b, resp.Status)
	}
}
