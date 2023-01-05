package clients

import "net/http"

type (
	Client interface {
		GetAccessToken(username string, password string) (string, error)
		RequestPayment() (bool, error)
	}

	Service struct {
		httpClient *http.Client
	}
)

func New() Service {
	return Service{
		httpClient: &http.Client{},
	}
}

func (s *Service) GetAccessToken(username string, password string) (string, error) {
	// make an api call to venmo api in go
	response, err := s.httpClient.Get("https://api.venmo.com/v1/oauth/access_token?client_id=123456&client_secret=123456&code")
	if err != nil {
		return "", err
	}

	return response.Status, nil
}
