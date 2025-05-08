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

type Constellation struct {
	ID       uint32   `json:"constellation_id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Position Position `json:"position,omitempty"`
	RegionID uint32   `json:"region_id,omitempty"`
	Systems  []uint32 `json:"systems,omitempty"`
}

type Region struct {
	ID             uint32   `json:"region_id,omitempty"`
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Constellations []uint32 `json:"constellations,omitempty"`
}

type Station struct {
	ID                       uint32   `json:"station_id,omitempty"`
	MaxDockableShipVolume    float32  `json:"max_dockable_ship_volume,omitempty"`
	Name                     string   `json:"name,omitempty"`
	OfficeRentalCost         float32  `json:"office_rental_cost,omitempty"`
	Owner                    uint32   `json:"owner,omitempty"`
	Position                 Position `json:"position,omitempty"`
	RaceID                   uint32   `json:"race_id,omitempty"`
	ReprocessingEfficiency   float32  `json:"reprocessing_efficiency,omitempty"`
	ReprocessingStationsTake float32  `json:"reprocessing_stations_take,omitempty"`
	Services                 []string `json:"services,omitempty"`
	SystemID                 uint32   `json:"system_id,omitempty"`
	TypeID                   uint32   `json:"type_id,omitempty"`
}

type Star struct {
	Age           uint64  `json:"age,omitempty"`
	Luminosity    float32 `json:"luminosity,omitempty"`
	Name          string  `json:"name,omitempty"`
	Radius        uint64  `json:"radius,omitempty"`
	SystemID      uint32  `json:"solar_system_id,omitempty"`
	SpectralClass string  `json:"spectral_class,omitempty"`
	Temperature   uint32  `json:"temperature,omitempty"`
	TypeID        uint32  `json:"type_id,omitempty"`
}

// NameRef is a reference to a name that is returned from esi
type NameRef struct {
	Category string `json:"category"`
	ID       uint   `json:"id"`
	Name     string `json:"name"`
}

// GetTypeIds get a list of all type ids in the game
func (esi Client) GetTypeIds() ([]uint32, error) {
	return esi.getIds("/v1/universe/types/")
}

// GetType gets the types information from esi
func (esi Client) GetType(id uint32) (UniverseType, error) {
	var item UniverseType
	err := esi.get(fmt.Sprintf("/v3/universe/types/%d/", id), &item)
	if err != nil {
		return UniverseType{}, err
	}

	return item, nil
}

func (esi Client) GetSystems() ([]uint32, error) {
	return esi.getIds("/latest/universe/systems/")
}

func (esi Client) GetConstellations() ([]uint32, error) {
	return esi.getIds("/latest/universe/constellations/")
}

func (esi Client) GetRegions() ([]uint32, error) {
	return esi.getIds("/latest/universe/regions/")
}

func (esi Client) GetSystem(id uint32) (SolarSystem, error) {
	var system SolarSystem
	err := esi.get(fmt.Sprintf("/latest/universe/systems/%d/", id), &system)
	if err != nil {
		return SolarSystem{}, err
	}

	return system, nil
}

func (esi Client) GetConstellation(id uint32) (Constellation, error) {
	var constellation Constellation
	err := esi.get(fmt.Sprintf("/latest/universe/constellations/%d/", id), &constellation)
	if err != nil {
		return Constellation{}, err
	}

	return constellation, nil
}

func (esi Client) GetRegion(id uint32) (Region, error) {
	var region Region
	err := esi.get(fmt.Sprintf("/latest/universe/regions/%d/", id), &region)
	if err != nil {
		return Region{}, err
	}

	return region, nil
}

func (esi Client) GetStargate(id uint32) (Stargate, error) {
	var gate Stargate
	err := esi.get(fmt.Sprintf("/latest/universe/stargates/%d/", id), &gate)
	if err != nil {
		return Stargate{}, err
	}

	return gate, nil
}

func (esi Client) GetStation(id uint32) (Station, error) {
	var station Station
	err := esi.get(fmt.Sprintf("/latest/universe/stations/%d/", id), &station)
	if err != nil {
		return Station{}, err
	}

	return station, nil
}

func (esi Client) GetStar(id uint32) (Star, error) {
	var star Star
	err := esi.get(fmt.Sprintf("/latest/universe/stars/%d/", id), &star)
	if err != nil {
		return Star{}, err
	}

	return star, nil
}

func (esi Client) GetNames(ids []uint) (map[uint]NameRef, error) {
	buffer, err := json.Marshal(ids)
	if err != nil {
		return nil, err
	}

	var names []NameRef
	err = esi.post("/v3/universe/names/", buffer, &names)
	if err != nil {
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
