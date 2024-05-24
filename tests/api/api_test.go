package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

const (
	baseurl = "http://127.0.0.1:9001/"
	geturl1 = baseurl + "api/v1/hello-world"
	geturl2 = baseurl + "api/v1/social/list"

	posturl     = baseurl + "api/v1/admin/login"
	numRequests = 100
	concurrency = 10
)

type KeyValue struct {
	Key   string
	Value string
}

func TestApi(t *testing.T) {
	//datas := []KeyValue{
	//	{"username", "super_admin"},
	//	{"password", "123456"},
	//}
	//testmain(arrgument(datas), postRequest)

	datas := []KeyValue{
		{"username", "aexliu"},
	}
	testmain(arrgument(datas), geturl2, getRequest)
	testmain(arrgument(datas), geturl1, getRequest)

}

func arrgument(data []KeyValue) map[string]string {
	var dataMap = make(map[string]string)
	for _, v := range data {
		dataMap[v.Key] = v.Value
	}
	return dataMap
}

func testmain(input map[string]string, link string, requestFunc func(map[string]string, string) (time.Duration, error)) {
	var wg sync.WaitGroup
	requestsChan := make(chan int, concurrency)
	resultsChan := make(chan time.Duration, numRequests)
	errorChan := make(chan struct{}, numRequests)

	// Worker function to send requests and measure latency
	worker := func(id int, requestFunc func() (time.Duration, error)) {
		for range requestsChan {
			latency, err := requestFunc()
			if err != nil {
				fmt.Printf("Worker %d: Request failed: %v\n", id, err)
				errorChan <- struct{}{}
				continue
			}
			resultsChan <- latency
		}
		wg.Done()
	}

	// Start workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(i, func() (time.Duration, error) { return requestFunc(input, link) })
	}

	// Send requests
	go func() {
		for i := 0; i < numRequests; i++ {
			requestsChan <- i
		}
		close(requestsChan)
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(resultsChan)
	close(errorChan)

	// Calculate results
	var totalLatency time.Duration
	var successfulRequests int
	var failedRequests int

	for latency := range resultsChan {
		totalLatency += latency
		successfulRequests++
	}

	for range errorChan {
		failedRequests++
	}

	avgLatency := totalLatency / time.Duration(successfulRequests)
	throughput := float64(successfulRequests) / totalLatency.Seconds()

	fmt.Printf("Total Requests: %d\n", numRequests)
	fmt.Printf("Successful Requests: %d\n", successfulRequests)
	fmt.Printf("Failed Requests: %d\n", failedRequests)
	fmt.Printf("Average Latency: %v\n", avgLatency)
	fmt.Printf("Throughput: %.2f requests/second\n", throughput)
}

func postRequest(params map[string]string, postUrl string) (time.Duration, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	start := time.Now()
	resp, err := client.Do(req)
	latency := time.Since(start)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status code: %v", resp.StatusCode)
	}

	return latency, nil
}

func getRequest(params map[string]string, getUrl string) (time.Duration, error) {
	query := getUrl + "?"
	for key, value := range params {
		query += key + "=" + url.QueryEscape(value) + "&"
	}

	start := time.Now()
	resp, err := http.Get(query)
	latency := time.Since(start)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status code: %v", resp.StatusCode)
	}

	return latency, nil
}
