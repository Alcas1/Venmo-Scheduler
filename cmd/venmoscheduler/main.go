package main

import (
	"fmt"

	venmo "github.com/Alcas1/Venmo-Scheduler/internal/clients/venmo"
)

func main() {
	venmoClient := venmo.New()
	accessToken, err := venmoClient.GetAccessToken("username", "password")
	fmt.Println(accessToken, err)
	// make an api call to venmo api in go
	// response, err := http.Get("https://api.venmo.com/v1/payments?access_token=123456")
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s	", err)
	// }

	// fmt.Println(response)
}
