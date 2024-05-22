package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

const (
	baseurl     = "http://127.0.0.1:9001/"
	geturl      = baseurl + "api/v1/hello-world"
	posturl     = baseurl + "api/v1/admin/login"
	numRequests = 1000
	concurrency = 10
)

func TestApi(t *testing.T) {
	testmain(generateRandomParams())
}

func generateRandomParams() map[string]interface{} {
	params := make(map[string]interface{})
	//paramCount := rand.Intn(5) + 1 // Random number of parameters between 1 and 5
	//
	//for i := 0; i < paramCount; i++ {
	//	key := fmt.Sprintf("param%d", i)
	//	value := fmt.Sprintf("value%d", rand.Intn(100))
	//	params[key] = value
	//}
	params["username"] = "super_admin"
	params["password"] = "123456"
	return params

}

func testmain(input map[string]interface{}) {
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
		go worker(i, func() (time.Duration, error) { return postRequest(input) })
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

func postRequest(params map[string]interface{}) (time.Duration, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
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

func getRequest(params map[string]interface{}) (time.Duration, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	start := time.Now()
	resp, err := http.Get(geturl + "?" + string(jsonData))
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
