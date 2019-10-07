package esi

import (
	"encoding/json"
	"fmt"
)

// MarketGroup is a group that appears on the market
type MarketGroup struct {
	Description   string   `json:"description,omitempty"`
	MarketGroupID uint32   `json:"market_group_id,omitempty"`
	Name          string   `json:"name,omitempty"`
	ParentGroupID uint32   `json:"parent_group_id,omitempty"`
	Types         []uint32 `json:"types,omitempty"`
}

// GetMarketGroupIds returns a list of all possible market group ids
func (esi Client) GetMarketGroupIds() ([]uint32, error) {
	body, error := esi.get("/latest/markets/groups/?datasource=tranquility")
	if error != nil {
		return nil, error
	}

	var groupIds []uint32
	if err := json.Unmarshal(body, &groupIds); err != nil {
		return nil, err
	}

	return groupIds, nil
}

// GetMarketGroup get the specified market group
func (esi Client) GetMarketGroup(id uint32) (*MarketGroup, error) {
	body, error := esi.get(fmt.Sprintf("/v1/markets/groups/%d/", id))
	if error != nil {
		return nil, error
	}

	var group MarketGroup
	if err := json.Unmarshal(body, &group); err != nil {
		return nil, err
	}

	return &group, error
}
