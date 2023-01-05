package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.venmo.com/v1/oauth/access_token?client_id=123456&client_secret=123456&code=123456")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s	", err)
	}

	fmt.Println(response)
}
