package main

import (
	"fmt"
	"net/http"
)

func main() {
	// make an api call to venmo api in go
	response, err := http.Get("https://api.venmo.com/v1/payments?access_token=123456")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s	", err)
	}

	fmt.Println(response)
}
