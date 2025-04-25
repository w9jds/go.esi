package esi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type dogmaAttributes struct {
	AttributeID uint32  `json:"attribute_id,omitempty"`
	Value       float64 `json:"value,omitempty"`
}

type dogmaEffects struct {
	EffectID  uint32 `json:"effect_id,omitempty"`
	IsDefault bool   `json:"is_default,omitempty"`
}

// UniverseType represents an item in eve online
type UniverseType struct {
	ID              uint32            `json:"type_id,omitempty"`
	Capacity        float32           `json:"capacity,omitempty"`
	Description     string            `json:"description,omitempty"`
	DogmaAttributes []dogmaAttributes `json:"dogma_attributes,omitempty"`
	DogmaEffects    []dogmaEffects    `json:"dogma_effects,omitempty"`
	GraphicID       uint32            `json:"graphic_id,omitempty"`
	GroupID         uint32            `json:"group_id,omitempty"`
	IconID          uint32            `json:"icon_id,omitempty"`
	MarketGroupID   uint32            `json:"market_group_id,omitempty"`
	Mass            float64           `json:"mass,omitempty"`
	Name            string            `json:"name,omitempty"`
	PackageVolume   float32           `json:"packaged_volume,omitempty"`
	PortionSize     uint32            `json:"portion_size,omitempty"`
	Published       bool              `json:"published,omitempty"`
	Radius          float32           `json:"radius,omitempty"`
	Volume          float32           `json:"volume,omitempty"`
}

type Planet struct {
	AstroidBelts []uint32 `json:"asteroid_belts,omitempty"`
	Moons        []uint32 `json:"moons,omitempty"`
	PlanetID     uint32   `json:"planet_id,omitempty"`
}

type SolarSystem struct {
	ID              uint32   `json:"system_id,omitempty"`
	ConstellationID uint32   `json:"constellation_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Planets         []Planet `json:"planets,omitempty"`
	Position        Position `json:"position,omitempty"`
	SecurityClass   string   `json:"security_class,omitempty"`
	SecurityStatus  float32  `json:"security_status,omitempty"`
	StarID          uint32   `json:"star_id,omitempty"`
	Stargates       []uint32 `json:"stargates,omitempty"`
	Stations        []uint32 `json:"stations,omitempty"`
}

type Stargate struct {
	ID          uint32   `json:"stargate_id,omitempty"`
	SystemID    uint32   `json:"system_id,omitempty"`
	Position    Position `json:"position,omitempty"`
	Name        string   `json:"name,omitempty"`
	TypeID      uint32   `json:"type_id,omitempty"`
	Destination struct {
		StargateID uint32 `json:"stargate_id,omitempty"`
		SystemID   uint32 `json:"system_id,omitempty"`
	} `json:"destination,omitempty"`
}

// NameRef is a reference to a name that is returned from esi
type NameRef struct {
	Category string `json:"category"`
	ID       uint   `json:"id"`
	Name     string `json:"name"`
}

// GetTypeIds get a list of all type ids in the game
func (esi Client) GetTypeIds() ([]uint32, error) {
	body, err := esi.get("/v1/universe/types/")
	if err != nil {
		return nil, err
	}

	var typeIds []uint32
	if err := json.Unmarshal(body, &typeIds); err != nil {
		return nil, err
	}

	return typeIds, nil
}

// GetType gets the types information from esi
func (esi Client) GetType(id uint32) (*UniverseType, error) {
	body, err := esi.get(fmt.Sprintf("/v3/universe/types/%d/", id))
	if err != nil {
		return nil, err
	}

	var item UniverseType
	if err := json.Unmarshal(body, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (esi Client) GetSystems() ([]uint32, error) {
	body, err := esi.get("/latest/universe/systems/")
	if err != nil {
		return nil, err
	}

	var systems []uint32
	if err := json.Unmarshal(body, &systems); err != nil {
		return nil, err
	}

	return systems, nil
}

func (esi Client) GetSystem(systemID uint32) (*SolarSystem, error) {
	body, err := esi.get(fmt.Sprintf("/latest/universe/systems/%d/", systemID))
	if err != nil {
		return nil, err
	}

	var system SolarSystem
	if err := json.Unmarshal(body, &system); err != nil {
		return nil, err
	}

	return &system, nil
}

func (esi Client) GetStargate(stargateId uint32) (*Stargate, error) {
	body, err := esi.get(fmt.Sprintf("/latest/universe/stargates/%d/", stargateId))
	if err != nil {
		return nil, err
	}

	var gate Stargate
	if err := json.Unmarshal(body, &gate); err != nil {
		return nil, err
	}

	return &gate, nil
}

// GetNames get a list of names from a list of ids
func (esi Client) GetNames(ids []uint) (map[uint]NameRef, error) {
	buffer, err := json.Marshal(ids)
	if err != nil {
		return nil, err
	}

	body, err := esi.post("/v3/universe/names/", buffer)
	if err != nil {
		return nil, err
	}

	var names []NameRef
	if err := json.Unmarshal(body, &names); err != nil {
		return nil, err
	}

	if len(ids) != len(names) {
		return nil, errors.New("names response didn't return same amount of items as original ids")
	}

	return mapNames(names), nil
}

func mapNames(names []NameRef) map[uint]NameRef {
	references := map[uint]NameRef{}

	for _, name := range names {
		references[name.ID] = name
	}

	return references
}
