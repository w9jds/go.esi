package esi

import (
	"fmt"
)

type FleetInfo struct {
  FleetBossID  int64 `json:"fleet_boss_id"`
  FleetID      int64 `json:"fleet_id"`
  Role         string `json:"role"`
  SquadID      int64 `json:"squad_id"`
  WingID       int64 `json:"wing_id"`
}

type FleetMember struct {
  CharacterID  int64 `json:"character_id"`
  JoinTime     string `json:"join_time"`
  Role         string `json:"role"`
  RoleName     string `json:"role_name"`
  ShipTypeID   int64 `json:"ship_type_id"`
  SolarSystemID int64 `json:"solar_system_id"`
  SquadID      int64 `json:"squad_id"`
  StationID    int64 `json:"station_id"`
  TakesFleetWarp bool  `json:"takes_fleet_warp"`
  WingID       int64 `json:"wing_id"`
}

type Squad struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type Wing struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Squads []Squad `json:"squads"`
}

func (esi Client) GetFleetInfo(characterID int64, token string) (FleetInfo, error) {
	var info FleetInfo
	err := esi.authGet(fmt.Sprintf("/characters/%d/online/", characterID), token, &info)
	if err != nil {
		return FleetInfo{}, err
	}

	return info, nil
}

func (esi Client) GetFleetMembers(fleetID int64, token string) ([]FleetMember, error) {
	var members []FleetMember
	err := esi.authGet(fmt.Sprintf("/fleets/%d/members/", fleetID), token, &members)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (esi Client) GetFleetWings(fleetID int64, token string) ([]Wing, error) {
	var wings []Wing
	err := esi.authGet(fmt.Sprintf("/fleets/%d/wings", fleetID), token, &wings)
	if err != nil {
		return nil, err
	}

	return wings, nil
}
