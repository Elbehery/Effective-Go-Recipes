package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Info struct {
	delay      time.Duration
	statusCode int
	url        string
	err        error
}

func siteInfo(url string) (Info, error) {
	//var info Info

	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return Info{}, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return Info{}, err
	}

	var info Info
	info.delay = time.Since(startTime)
	info.statusCode = resp.StatusCode
	return info, nil
}

func sitesInfo(urls []string) (map[string]Info, error) {
	out := make(map[string]Info)
	for _, url := range urls {
		info, err := siteInfo(url)
		if err != nil {
			return nil, err
		}
		out[url] = info
	}
	return out, nil
}

func siteInfosFix(urls []string) (map[string]Info, error) {
	ch := make(chan Info)
	defer close(ch)

	for _, url := range urls {
		go func(u string) {
			info, err := siteInfo(u)
			info.url = u
			info.err = err
			ch <- info
		}(url)
	}

	out := make(map[string]Info)

	for range urls {
		info := <-ch
		if info.err != nil {
			return nil, info.err
		}
		out[info.url] = info
	}

	return out, nil
}

func main() {
	start := time.Now()

	urls := []string{
		"https://www.apple.com/",
		"https://www.microsoft.com/",
		"https://www.ibm.com/",
		"https://www.dell.com/",
	}

	infos, err := siteInfosFix(urls)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	for url, info := range infos {
		fmt.Printf("%s: %+v\n", url, info)
	}

	duration := time.Since(start)
	fmt.Printf("%d sites in %v\n", len(urls), duration)
}
