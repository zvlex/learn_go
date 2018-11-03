package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	urlPrefix = "http://"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	var fullURL string

	start := time.Now()

	if !strings.HasPrefix(url, urlPrefix) {
		fullURL = urlPrefix + url
	}

	resp, err := http.Get(fullURL)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	file, err := os.Create(url + ".html")

	if err != nil {
		fmt.Printf("Can't create file")
	}

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		fmt.Printf("Can't write in file")
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
