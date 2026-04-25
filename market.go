package esi

import (
	"fmt"
)

// MarketGroup is a group that appears on the market
type MarketGroup struct {
	Description   string   `json:"description,omitempty"`
	MarketGroupID int64   `json:"market_group_id,omitempty"`
	Name          string   `json:"name,omitempty"`
	ParentGroupID int64   `json:"parent_group_id,omitempty"`
	Types         []int64 `json:"types,omitempty"`
}

// GetMarketGroupIds returns a list of all possible market group ids
func (esi Client) GetMarketGroupIds() ([]int64, error) {
	return esi.getIds("/markets/groups/")
}

// GetMarketGroup get the specified market group
func (esi Client) GetMarketGroup(id int64) (*MarketGroup, error) {
	var group MarketGroup
	error := esi.get(fmt.Sprintf("/markets/groups/%d/", id), &group)
	if error != nil {
		return nil, error
	}

	return &group, error
}
