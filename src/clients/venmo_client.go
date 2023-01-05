package clients

import "net/http"

type VenmoClient interface {
	GetAccessToken() (*http.Response, error)
	RequestPayment() (*http.Response, error)
}
