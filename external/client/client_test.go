package external

import (
	"fmt"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	//client := NewClient("https://www.oklink.com/api/v5/explorer/blockchain", "baeffec3-4eb6-4346-9dc3-f1bc8f4345cd")
	//
	//endpoint := "/summary?chainShortName=ETH"
	//method := "GET"
	//params := map[string]string{}
	//
	//response, err := client.SendRequest(method, endpoint, params)
	//if err != nil {
	//	log.Fatalf("Request failed: %v", err)
	//}
	//
	//fmt.Printf("Response: %s\n", response)
	for i := 0; i < 100; i++ {
		test()
	}

}

func test() {

	client := NewClient("http://127.0.0.1:9001/api/v1/social/list", "")

	endpoint := ""
	method := "GET"
	params := map[string]string{
		"username": "aexliu",
	}

	response, err := client.SendRequest(method, endpoint, params)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	fmt.Printf("Response: %s\n", response)
}
