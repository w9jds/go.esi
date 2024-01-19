package esi

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Client is a client for communication with the eve online api
type Client struct {
	baseURI string
	client  *http.Client
}

const baseURI = "https://esi.evetech.net"

// CreateClient creates a new instance of the Client
func CreateClient(httpClient *http.Client) *Client {
	return &Client{
		baseURI: baseURI,
		client:  httpClient,
	}
}

func attachHeaders(request *http.Request) *http.Request {
	request.Header.Add("Accept", "application/json")
	return request
}

func authHeader(request *http.Request, token string) *http.Request {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return request
}

func (esi Client) get(path string) ([]byte, error) {
	request, error := http.NewRequest("GET", baseURI+path, nil)
	if error != nil {
		return nil, error
	}
	return esi.do(attachHeaders(request))
}

func (esi Client) authGet(path string, token string) ([]byte, error) {
	request, err := http.NewRequest("GET", baseURI+path, nil)
	if err != nil {
		return nil, err
	}

	request = authHeader(request, token)
	return esi.do(attachHeaders(request))
}

func (esi Client) post(path string, content []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", baseURI+path, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}

	return esi.do(attachHeaders(request))
}

func (esi Client) do(request *http.Request) ([]byte, error) {
	var content []byte
	var err error

	for i := 0; i < 3; i++ {
		delay := 5 * time.Second

		response, error := esi.client.Do(request)
		if error != nil {
			log.Printf("%v\n", error)
			continue
		} else if response.StatusCode < 200 || response.StatusCode > 299 {
			content, err = io.ReadAll(response.Body)

			if response.StatusCode == 420 {
				// on error limited wait 60 seconds before proceeding
				log.Printf("I've been rate limited, waiting 60 seconds before trying again\n")
				delay = 1 * time.Minute
			}

			// Don't bother retrying three times when you don't have permissions to make the request in the first place
			if response.StatusCode == 403 || response.StatusCode == 401 {
				return nil, fmt.Errorf("%v", map[string]interface{}{
					"error":   fmt.Sprintf("Status %d: Unauthorized", response.StatusCode),
					"url":     request.URL,
					"status":  response.StatusCode,
					"content": string(content),
					"caught":  err,
				})
			}

			if response.StatusCode == 400 {
				return nil, fmt.Errorf("%v", map[string]interface{}{
					"error":   "Bad Request",
					"url":     request.URL,
					"status":  response.StatusCode,
					"content": string(content),
					"caught":  err,
				})
			}

			if response.StatusCode == 404 || strings.Contains(string(content), "404") {
				return nil, fmt.Errorf("%v", map[string]interface{}{
					"error":   "Status 404: Not Found",
					"url":     request.URL,
					"status":  response.StatusCode,
					"content": string(content),
					"caught":  err,
				})
			}

			if err != nil {
				log.Printf("%v\n", err)
				time.Sleep(delay)
				continue
			} else {
				time.Sleep(delay)
				continue
			}
		} else {
			return io.ReadAll(response.Body)
		}
	}

	var errorPayload = map[string]interface{}{
		"error":   "Failed to make request after 3 attempts",
		"url":     request.URL,
		"content": string(content),
		"caught":  err,
	}

	return nil, fmt.Errorf("%v", errorPayload)
}
