package esi

import (
	"encoding/json"
	"fmt"
)

// CharacterDetails references information from the character endpoint
type CharacterDetails struct {
	AllianceID     uint32  `json:"alliance_id,omitempty"`
	Birthday       string  `json:"birthday"`
	BloodlineID    uint32  `json:"bloodline_id,omitempty"`
	CorporationID  uint32  `json:"corporation_id"`
	Description    string  `json:"description,omitempty"`
	FactionID      uint32  `json:"faction_id,omitempty"`
	Gender         string  `json:"gender"`
	Name           string  `json:"name"`
	RaceID         uint32  `json:"race_id"`
	SecurityStatus float32 `json:"security_status"`
	Title          string  `json:"title"`
}

// OnlineStatus references information about the last login for a character
type OnlineStatus struct {
	LastLogin  string `json:"last_login,omitempty"`
	LastLogout string `json:"last_logout,omitempty"`
	Logins     uint   `json:"logins,omitempty"`
	Online     bool   `json:"online"`
}

// Location reference to the location the character is currently located
type Location struct {
	SolarSystemID uint `json:"solar_system_id"`
	StationID     uint `json:"station_id,omitempty"`
	StructureID   uint `json:"structure_id,omitempty"`
}

// Ship entity is the ship that the character is currently in
type Ship struct {
	ShipItemID uint   `json:"ship_item_id"`
	ShipName   string `json:"ship_name"`
	ShipTypeID uint   `json:"ship_type_id"`
}

// Affiliation represents a characters corp affiliations
type Affiliation struct {
	AllianceID  uint32 `json:"alliance_id,omitempty"`
	CharacterID uint32 `json:"character_id,omitempty"`
	CorpID      uint32 `json:"corporation_id,omitempty"`
	FactionID   uint32 `json:"faction_id,omitempty"`
}

// Roles represents the roles a character currently has
type Roles struct {
	Roles      []string `json:"roles,omitempty"`
	BaseRoles  []string `json:"roles_at_base,omitempty"`
	HQRoles    []string `json:"roles_at_hq,omitempty"`
	OtherRoles []string `json:"roles_at_other,omitempty"`
}

// Title represents a title
type Title struct {
	Name string `json:"name,omitempty"`
	ID   uint32 `json:"title_id,omitempty"`
}

// CorporationHistory is a history record for a corp the character belonged to
type CorporationHistory struct {
	ID        uint32 `json:"corporation_id,omitempty"`
	Deleted   bool   `json:"is_deleted,omitempty"`
	RecordID  uint32 `json:"record_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
}

func (esi Client) GetCharacterCorpHistory(characterID uint32) ([]CorporationHistory, error) {
	var history []CorporationHistory
	err := esi.get(fmt.Sprintf("/v2/characters/%d/corporationhistory/", characterID), &history)
	if err != nil {
		return []CorporationHistory{}, err
	}

	return history, nil
}

// IsCharacterOnline gets if the character is currently online
func (esi Client) IsCharacterOnline(characterID uint32, token string) (OnlineStatus, error) {
	var status OnlineStatus
	err := esi.authGet(fmt.Sprintf("/v3/characters/%d/online/", characterID), token, &status)
	if err != nil {
		return OnlineStatus{}, err
	}

	return status, nil
}

// GetCharacterLocation get the character's current location
func (esi Client) GetCharacterLocation(characterID uint32, token string) (Location, error) {
	var location Location
	err := esi.authGet(fmt.Sprintf("/v2/characters/%d/location/", characterID), token, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

// GetCharacterShip get the character's current ship
func (esi Client) GetCharacterShip(characterID uint32, token string) (Ship, error) {
	var ship Ship
	err := esi.authGet(fmt.Sprintf("/v2/characters/%d/ship/", characterID), token, &ship)
	if err != nil {
		return Ship{}, err
	}

	return ship, nil
}

// GetCharacterRoles gets the current for this character
func (esi Client) GetCharacterRoles(characterID uint32, token string) (Roles, error) {
	var roles Roles
	err := esi.authGet(fmt.Sprintf("/v3/characters/%d/roles/", characterID), token, &roles)
	if err != nil {
		return Roles{}, err
	}

	return roles, nil
}

// GetCharacterTitles returns a list of a characters awarded titles
func (esi Client) GetCharacterTitles(characterID uint32, token string) ([]Title, error) {
	var titles []Title
	error := esi.authGet(fmt.Sprintf("/v2/characters/%d/titles/", characterID), token, &titles)
	if error != nil {
		return nil, error
	}

	return titles, nil
}

// GetCharacterDetails retrieves the characters basic information from the characterID
func (esi Client) GetCharacterDetails(characterID uint32) (*CharacterDetails, error) {
	var details CharacterDetails
	err := esi.get(fmt.Sprintf("/v5/characters/%d/", characterID), &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// GetCharacterAffiliations get the affiliations of all passed of characterIds
func (esi Client) GetCharacterAffiliations(ids []uint32) ([]Affiliation, error) {
	buffer, err := json.Marshal(ids)
	if err != nil {
		return nil, err
	}

	var affiliations []Affiliation
	err = esi.post("/v2/characters/affiliation/", buffer, &affiliations)
	if err != nil {
		return nil, err
	}

	return affiliations, nil
}
