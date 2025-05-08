package esi

type ESIStatus struct {
	Players   uint   `json:"players"`
	Version   string `json:"server_version"`
	StartTime string `json:"start_time"`
	VIP       bool   `json:"vip,omitempty"`
}

// GetServerStatus get the status of the ESI cluster
func (esi Client) GetServerStatus() (*ESIStatus, error) {
	var status ESIStatus
	err := esi.get("/v2/status", &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}
