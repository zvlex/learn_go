package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	urlPrefix = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, urlPrefix) {
			url = urlPrefix + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", resp.Status)
	}
}
