package esi

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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
	request, error := http.NewRequest("POST", baseURI+path, bytes.NewBuffer(content))
	if error != nil {
		return nil, error
	}

	return esi.do(attachHeaders(request))
}

func (esi Client) do(request *http.Request) ([]byte, error) {

	for i := 0; i < 3; i++ {
		delay := 5 * time.Second

		response, error := esi.client.Do(request)
		if error != nil {
			log.Println(error)
			continue
		} else if response.StatusCode < 200 || response.StatusCode > 299 {

			// Don't bother retrying three times when you don't have permissions to make the request in the first place
			if response.StatusCode == 403 || response.StatusCode == 401 {
				log.Printf("Status %d: Unauthorized\n", response.StatusCode)
				break
			}

			message, error := io.ReadAll(response.Body)

			if response.StatusCode == 420 {
				// on error limited wait 60 seconds before proceeding
				delay = 1 * time.Minute
			}

			if error != nil {
				log.Println(error)
				time.Sleep(delay)
				continue
			} else {
				log.Println(string(message))
				time.Sleep(delay)
				continue
			}
		} else {
			return io.ReadAll(response.Body)
		}
	}

	return nil, errors.New("failed esi requests 3 times, gave up")
}
