package esi

import (
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
	return esi.getIds("/latest/markets/groups/")
}

// GetMarketGroup get the specified market group
func (esi Client) GetMarketGroup(id uint32) (*MarketGroup, error) {
	var group MarketGroup
	error := esi.get(fmt.Sprintf("/v1/markets/groups/%d/", id), &group)
	if error != nil {
		return nil, error
	}

	return &group, error
}
