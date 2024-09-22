package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type siteConfig struct {
	Url             string
	AcceptableCodes []int
	Frequency       time.Duration
}

type Result struct {
	Url        string
	StatusCode int
	Up         bool
}

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type DefaultClient struct{}

func (c *DefaultClient) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func checker(config siteConfig, client HttpClient, results chan<- Result) {
	result := Result{Url: config.Url, StatusCode: -1, Up: false}
	resp, err := client.Get(config.Url)
	if err != nil {
		results <- result
		return
	}
	defer resp.Body.Close()
	result.StatusCode = resp.StatusCode
	for _, code := range config.AcceptableCodes {
		if resp.StatusCode == code {
			result.Up = true
			break
		}
	}
	results <- result
}

func scheduleCheck(config siteConfig, client HttpClient, results chan<- Result) {
	go func() {
		ticker := time.NewTicker(config.Frequency)
		for {
			select {
			case <-ticker.C:
				checker(config, client, results)
			}
		}
	}()
}

func main() {
	sites := []siteConfig{
		{"https://google.com", []int{200, 201}, time.Second * 1},
		{"http://localhost:1983", []int{200, 201}, time.Second * 3},
		{"https://go.dev", []int{200, 201}, time.Second * 2},
	}
	results := make(chan Result)
	client := &DefaultClient{}
	for _, site := range sites {
		scheduleCheck(site, client, results)
	}

	file, err := os.Create("log1.txt")
	if err != nil {
		log.Fatal(err)
	}
	for res := range results {
		if res.Up {
			fmt.Fprintf(os.Stdout, "%s is up and running with status code %d\n", res.Url, res.StatusCode)
			file.WriteString(fmt.Sprintf("%s is up and running with status code %d", res.Url, res.StatusCode) + "\n")
		} else {
			fmt.Fprintf(os.Stdout, "%s is down with status code %d\n", res.Url, res.StatusCode)
			file.WriteString(fmt.Sprintf("%s is up and running with status code %d", res.Url, res.StatusCode) + "\n")
		}
	}
}
