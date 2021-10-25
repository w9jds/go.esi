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

// IsCharacterOnline gets if the character is currently online
func (esi Client) IsCharacterOnline(characterID uint32, token string) (*OnlineStatus, error) {
	body, error := esi.authGet(fmt.Sprintf("/v3/characters/%d/online/", characterID), token)
	if error != nil {
		return nil, error
	}

	var status OnlineStatus
	error = json.Unmarshal(body, &status)
	if error != nil {
		return nil, error
	}

	return &status, nil
}

// GetCharacterLocation get the character's current location
func (esi Client) GetCharacterLocation(characterID uint32, token string) (*Location, error) {
	body, error := esi.authGet(fmt.Sprintf("/v2/characters/%d/location/", characterID), token)
	if error != nil {
		return nil, error
	}

	var location Location
	error = json.Unmarshal(body, &location)
	if error != nil {
		return nil, error
	}

	return &location, nil
}

// GetCharacterShip get the character's current ship
func (esi Client) GetCharacterShip(characterID uint32, token string) (*Ship, error) {
	body, error := esi.authGet(fmt.Sprintf("/v2/characters/%d/ship/", characterID), token)
	if error != nil {
		return nil, error
	}

	var ship Ship
	error = json.Unmarshal(body, &ship)
	if error != nil {
		return nil, error
	}

	return &ship, nil
}

// GetCharacterRoles gets the current for this character
func (esi Client) GetCharacterRoles(characterID uint32, token string) (*Roles, error) {
	body, error := esi.authGet(fmt.Sprintf("/v3/characters/%d/roles/", characterID), token)
	if error != nil {
		return nil, error
	}

	var roles Roles
	error = json.Unmarshal(body, &roles)
	if error != nil {
		return nil, error
	}

	return &roles, nil
}

// GetCharacterTitles returns a list of a characters awarded titles
func (esi Client) GetCharacterTitles(characterID uint32, token string) ([]Title, error) {
	body, error := esi.authGet(fmt.Sprintf("/v2/characters/%d/titles/", characterID), token)
	if error != nil {
		return nil, error
	}

	var titles []Title
	if err := json.Unmarshal(body, &titles); err != nil {
		return nil, err
	}

	return titles, nil
}

// GetCharacterDetails retrieves the characters basic information from the characterID
func (esi Client) GetCharacterDetails(characterID uint32) (*CharacterDetails, error) {
	body, err := esi.get(fmt.Sprintf("/v5/characters/%d/", characterID))
	if err != nil {
		return nil, err
	}

	var details CharacterDetails
	if err := json.Unmarshal(body, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

// GetCharacterAffiliations get the affiliations of all passed of characterIds
func (esi Client) GetCharacterAffiliations(ids []uint32) ([]Affiliation, error) {
	buffer, error := json.Marshal(ids)
	if error != nil {
		return nil, error
	}

	body, error := esi.post("/v2/characters/affiliation/", buffer)
	if error != nil {
		return nil, error
	}

	var affiliations []Affiliation
	if err := json.Unmarshal(body, &affiliations); err != nil {
		return nil, err
	}

	return affiliations, nil
}
