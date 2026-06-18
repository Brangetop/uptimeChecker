package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkWebsite(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	initTime := time.Now()

	resp, err := client.Head(url)
	if err != nil {
		fmt.Println(url, "ERROR:", err)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(initTime)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println(url, "OK", "Duration: ", duration)
	} else {
		fmt.Println(url, "FAIL", "StatusCode: ", resp.StatusCode)
	}
}

func main() {
	websitePool := []string{
		"https://google.com/",
		"https://go.dev/",
		"https://github.com",
		"https://youtube.com/",
		"https://wikipedia.org/",
	}

	var wg sync.WaitGroup

	for _, url := range websitePool {
		wg.Add(1)
		go checkWebsite(url, &wg)
	}
	wg.Wait()
}
