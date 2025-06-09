package esi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"k8s.io/klog"
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

func (esi Client) get(path string, result interface{}) error {
	request, err := http.NewRequest("GET", baseURI+path, nil)
	if err != nil {
		return err
	}

	data, err := esi.do(attachHeaders(request))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	return nil
}

func (esi Client) authGet(path string, token string, result interface{}) error {
	request, err := http.NewRequest("GET", baseURI+path, nil)
	if err != nil {
		return err
	}

	request = authHeader(request, token)
	data, err := esi.do(attachHeaders(request))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	return nil
}

func (esi Client) post(path string, content []byte, result interface{}) error {
	request, err := http.NewRequest("POST", baseURI+path, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	data, err := esi.do(attachHeaders(request))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	return nil
}

type ResponseError struct {
	Path       string
	StatusCode int
	Error      error
}

func (esi Client) do(request *http.Request) ([]byte, error) {
	for i := 0; i < 3; i++ {
		delay := 5 * time.Second

		response, error := esi.client.Do(request)
		if error != nil {
			klog.Error(ResponseError{
				Path:  request.URL.Path,
				Error: error,
			})
			time.Sleep(delay)
			continue
		} else if response.StatusCode < 200 || response.StatusCode > 299 {
			log := ResponseError{
				Path:       request.URL.Path,
				StatusCode: response.StatusCode,
			}

			// Don't bother retrying three times when you don't have permissions to make the request in the first place
			if response.StatusCode == 403 || response.StatusCode == 401 {
				klog.Error(log)
				break
			}

			message, error := io.ReadAll(response.Body)

			// Don't bother retrying three times when rate limited
			if response.StatusCode == 420 || response.StatusCode == 404 {
				log.Error = fmt.Errorf("%s", message)
				klog.Error(log)
				break
			}

			if error != nil {
				log.Error = error
			} else {
				log.Error = fmt.Errorf("%s", message)
			}

			klog.Error(log)
			time.Sleep(delay)
			continue
		} else {
			return io.ReadAll(response.Body)
		}
	}

	return nil, fmt.Errorf("Failed Request %s After 3 Tries", request.URL.Path)
}

func (esi Client) getIds(path string) ([]uint32, error) {
	var ids []uint32
	err := esi.get(path, &ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
