package esi

import (
	"fmt"
)

// KillMail that is recieved from eve online
type KillMail struct {
	ID        int64     `json:"killmail_id,omitempty"`
	Time      string     `json:"killmail_time,omitempty"`
	SystemID  int64     `json:"solar_system_id,omitempty"`
	Victim    victim     `json:"victim,omitempty"`
	Attackers []attacker `json:"attackers,omitempty"`
	WarID     int64     `json:"war_id,omitempty"`
}

type victim struct {
	ID            int64     `json:"character_id,omitempty"`
	AllianceID    int64     `json:"alliance_id,omitempty"`
	CorporationID int64     `json:"corporation_id,omitempty"`
	DamageTaken   uint64     `json:"damage_taken,omitempty"`
	Items         []KillItem `json:"items,omitempty"`
	ShipTypeID    int64     `json:"ship_type_id,omitempty"`
	Position      Position   `json:"position,omitempty"`
}

type Position struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}

type attacker struct {
	ID             int64  `json:"character_id,omitempty"`
	AllianceID     int64  `json:"alliance_id,omitempty"`
	CorporationID  int64  `json:"corporation_id,omitempty"`
	DamageDone     uint64  `json:"damage_done,omitempty"`
	FinalBlow      bool    `json:"final_blow,omitempty"`
	SecurityStatus float32 `json:"security_status,omitempty"`
	ShipTypeID     int64  `json:"ship_type_id,omitempty"`
	WeaponTypeID   int64  `json:"weapon_type_id,omitempty"`
}

// KillItem is an item that is found on a killmail
type KillItem struct {
	ID                int64 `json:"item_type_id,omitempty"`
	Flag              int16  `json:"flag,omitempty"`
	QuantityDropped   int64 `json:"quantity_dropped,omitempty"`
	QuantityDestroyed int64 `json:"quantity_destroyed,omitempty"`
	Singleton         int8   `json:"singleton,omitempty"`
}

// KillFitting the fitting built from the items on the victim's killmail
type KillFitting struct {
	SubSystemSlot map[int64]*KillItem
	HighSlot      map[int64]*KillItem
	MedSlot       map[int64]*KillItem
	LoSlot        map[int64]*KillItem
	RigSlot       map[int64]*KillItem
	FighterBay    map[int64]*KillItem
	ServiceSlot   map[int64]*KillItem
	Cargo         map[int64]*KillItem
	DroneBay      map[int64]*KillItem
}

// GetKillMail retrieves a specific killmail from ESI
func (esi Client) GetKillMail(killID int64, hash string, withFitting bool) (*KillMail, *KillFitting, error) {
	var killmail KillMail
	err := esi.get(fmt.Sprintf("/killmails/%d/%s/", killID, hash), &killmail)
	if err != nil {
		return nil, nil, err
	}

	if withFitting {
		return &killmail, killmail.BuildShipFitting(), nil
	}

	return &killmail, nil, nil
}

func updateFittingItem(group map[int64]*KillItem, item KillItem) {
	if current, ok := group[item.ID]; ok {
		current.QuantityDestroyed += item.QuantityDestroyed
		current.QuantityDropped += item.QuantityDropped
	} else {
		group[item.ID] = &item
	}
}

func (killmail KillMail) BuildShipFitting() *KillFitting {
	fit := &KillFitting{
		SubSystemSlot: map[int64]*KillItem{},
		HighSlot:      map[int64]*KillItem{},
		MedSlot:       map[int64]*KillItem{},
		LoSlot:        map[int64]*KillItem{},
		RigSlot:       map[int64]*KillItem{},
		Cargo:         map[int64]*KillItem{},
		DroneBay:      map[int64]*KillItem{},
	}

	for _, item := range killmail.Victim.Items {
		if item.Flag == 5 {
			updateFittingItem(fit.Cargo, item)
		} else if item.Flag == 87 {
			updateFittingItem(fit.DroneBay, item)
		} else if item.Flag >= 27 && item.Flag <= 34 {
			updateFittingItem(fit.HighSlot, item)
		} else if item.Flag >= 19 && item.Flag <= 26 {
			updateFittingItem(fit.MedSlot, item)
		} else if item.Flag >= 11 && item.Flag <= 18 {
			updateFittingItem(fit.LoSlot, item)
		} else if item.Flag >= 92 && item.Flag <= 94 {
			updateFittingItem(fit.RigSlot, item)
		} else if item.Flag >= 125 && item.Flag <= 128 {
			updateFittingItem(fit.SubSystemSlot, item)
		}
	}

	return fit
}
