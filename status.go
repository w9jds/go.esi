package esi

import "encoding/json"

type ESIStatus struct {
	Players   uint   `json:"players"`
	Version   string `json:"server_version"`
	StartTime string `json:"start_time"`
	VIP       bool   `json:"vip,omitempty"`
}

// GetServerStatus get the status of the ESI cluster
func (esi Client) GetServerStatus() (*ESIStatus, error) {
	body, err := esi.get("/v2/status")
	if err != nil {
		return nil, err
	}

	var status ESIStatus
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}

	return &status, nil
}
