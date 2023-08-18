package esi

import (
	"encoding/json"
	"fmt"
)

// MarketGroup is a group that appears on the market
type MarketGroup struct {
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	MarketGroupID uint32   `json:"market_group_id,omitempty"`
	ParentGroupID uint32   `json:"parent_group_id,omitempty"`
	Types         []uint32 `json:"types,omitempty"`
}

type MarketPrice struct {
	TypeID        int64   `json:"type_id,omitempty"`
	AveragePrice  float64 `json:"average_price,omitempty"`
	AdjustedPrice float64 `json:"adjusted_price,omitempty"`
}

type MarketOrder struct {
	ID            int64   `json:"order_id,omitempty"`
	Type          int64   `json:"type_id,omitempty"`
	StationID     int64   `json:"location_id,omitempty"`
	SystemID      int64   `json:"system_id,omitempty"`
	Volume        int64   `json:"volume_remain,omitempty"`
	MinVolume     int64   `json:"min_volume,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Buy           bool    `json:"is_buy_order,omitempty"`
	Duration      int64   `json:"duration,omitempty"`
	Issued        string  `json:"issued,omitempty"`
	VolumeEntered int64   `json:"volumeEntered,omitempty"`
	Range         string  `json:"range,omitempty"`
}

// GetMarketGroupIds returns a list of all possible market group ids
func (esi Client) GetMarketGroupIds() ([]uint32, error) {
	body, _, error := esi.get("/latest/markets/groups/?datasource=tranquility")
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
	body, _, error := esi.get(fmt.Sprintf("/v1/markets/groups/%d/", id))
	if error != nil {
		return nil, error
	}

	var group MarketGroup
	if err := json.Unmarshal(body, &group); err != nil {
		return nil, err
	}

	return &group, error
}

func (esi Client) GetMarketPrices() (*[]MarketPrice, error) {
	body, _, error := esi.get("/v1/markets/prices/")
	if error != nil {
		return nil, error
	}

	var prices []MarketPrice
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, err
	}

	return &prices, error
}

func (esi Client) GetMarketOrders(page int32, orderType string, regionID int32) (*[]MarketOrder, *Page, error) {
	body, headers, error := esi.get(fmt.Sprintf("/v1/markets/%d/orders/?datasource=tranquility&order_type=%s&page=%d", regionID, orderType, page))
	if error != nil {
		return nil, nil, error
	}

	var orders []MarketOrder
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, nil, err
	}

	return &orders, getPage(page, headers), error
}
