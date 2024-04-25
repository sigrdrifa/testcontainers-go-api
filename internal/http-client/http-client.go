package http_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AgeResponse struct {
	Count uint64
	Name  string
	Age   int
}

type client struct {
	baseUrl    string
	httpClient *http.Client
}

// Creates a new Client instance
func NewClient(baseUrl string) (*client, error) {
	if baseUrl == "" {
		return nil, errors.New("Invalid baseUrl provided")
	}

	httpClient := http.Client{Timeout: time.Duration(5) * time.Second}
	return &client{
		baseUrl:    baseUrl,
		httpClient: &httpClient,
	}, nil
}

func (c *client) GetAge(name string) (*AgeResponse, error) {

	resp, err := c.httpClient.Get(
		fmt.Sprintf("%s?name=%s", c.baseUrl, name))
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body : %s\n", body)
	ageResponse := AgeResponse{}

	err = json.Unmarshal(body, &ageResponse)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	return &ageResponse, nil
}
