package esi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type dogmaAttributes struct {
	AttributeID int64  `json:"attribute_id,omitempty"`
	Value       float64 `json:"value,omitempty"`
}

type dogmaEffects struct {
	EffectID  int64 `json:"effect_id,omitempty"`
	IsDefault bool   `json:"is_default,omitempty"`
}

// UniverseType represents an item in eve online
type UniverseType struct {
	ID              int64            `json:"type_id,omitempty"`
	Capacity        float32           `json:"capacity,omitempty"`
	Description     string            `json:"description,omitempty"`
	DogmaAttributes []dogmaAttributes `json:"dogma_attributes,omitempty"`
	DogmaEffects    []dogmaEffects    `json:"dogma_effects,omitempty"`
	GraphicID       int64            `json:"graphic_id,omitempty"`
	GroupID         int64            `json:"group_id,omitempty"`
	IconID          int64            `json:"icon_id,omitempty"`
	MarketGroupID   int64            `json:"market_group_id,omitempty"`
	Mass            float64           `json:"mass,omitempty"`
	Name            string            `json:"name,omitempty"`
	PackageVolume   float32           `json:"packaged_volume,omitempty"`
	PortionSize     int64            `json:"portion_size,omitempty"`
	Published       bool              `json:"published,omitempty"`
	Radius          float32           `json:"radius,omitempty"`
	Volume          float32           `json:"volume,omitempty"`
}

type GroupInfo struct {
	ID int64 `json:"group_id,omitempty"`
	CategoryID int64 `json:"category_id,omitempty"`
	Name string `json:"name,omitempty"`
	Published bool `json:"published,omitempty"`
	Types []int64 `json:"types,omitempty"`
}

type Planet struct {
	AstroidBelts []int64 `json:"asteroid_belts,omitempty"`
	Moons        []int64 `json:"moons,omitempty"`
	PlanetID     int64   `json:"planet_id,omitempty"`
}

type SolarSystem struct {
	ID              int64   `json:"system_id,omitempty"`
	ConstellationID int64   `json:"constellation_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Planets         []Planet `json:"planets,omitempty"`
	Position        Position `json:"position,omitempty"`
	SecurityClass   string   `json:"security_class,omitempty"`
	SecurityStatus  float32  `json:"security_status,omitempty"`
	StarID          int64   `json:"star_id,omitempty"`
	Stargates       []int64 `json:"stargates,omitempty"`
	Stations        []int64 `json:"stations,omitempty"`
}

type Stargate struct {
	ID          int64   `json:"stargate_id,omitempty"`
	SystemID    int64   `json:"system_id,omitempty"`
	Position    Position `json:"position,omitempty"`
	Name        string   `json:"name,omitempty"`
	TypeID      int64   `json:"type_id,omitempty"`
	Destination struct {
		StargateID int64 `json:"stargate_id,omitempty"`
		SystemID   int64 `json:"system_id,omitempty"`
	} `json:"destination,omitempty"`
}

type Constellation struct {
	ID       int64   `json:"constellation_id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Position Position `json:"position,omitempty"`
	RegionID int64   `json:"region_id,omitempty"`
	Systems  []int64 `json:"systems,omitempty"`
}

type Region struct {
	ID             int64   `json:"region_id,omitempty"`
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Constellations []int64 `json:"constellations,omitempty"`
}

type Station struct {
	ID                       int64   `json:"station_id,omitempty"`
	MaxDockableShipVolume    float32  `json:"max_dockable_ship_volume,omitempty"`
	Name                     string   `json:"name,omitempty"`
	OfficeRentalCost         float32  `json:"office_rental_cost,omitempty"`
	Owner                    int64   `json:"owner,omitempty"`
	Position                 Position `json:"position,omitempty"`
	RaceID                   int64   `json:"race_id,omitempty"`
	ReprocessingEfficiency   float32  `json:"reprocessing_efficiency,omitempty"`
	ReprocessingStationsTake float32  `json:"reprocessing_stations_take,omitempty"`
	Services                 []string `json:"services,omitempty"`
	SystemID                 int64   `json:"system_id,omitempty"`
	TypeID                   int64   `json:"type_id,omitempty"`
}

type Star struct {
	Age           uint64  `json:"age,omitempty"`
	Luminosity    float32 `json:"luminosity,omitempty"`
	Name          string  `json:"name,omitempty"`
	Radius        uint64  `json:"radius,omitempty"`
	SystemID      int64  `json:"solar_system_id,omitempty"`
	SpectralClass string  `json:"spectral_class,omitempty"`
	Temperature   int64  `json:"temperature,omitempty"`
	TypeID        int64  `json:"type_id,omitempty"`
}

// NameRef is a reference to a name that is returned from esi
type NameRef struct {
	Category string `json:"category"`
	ID       uint   `json:"id"`
	Name     string `json:"name"`
}

// GetTypeIds get a list of all type ids in the game
func (esi Client) GetTypeIds() ([]int64, error) {
	return esi.getIds("/universe/types/")
}

// GetType gets the types information from esi
func (esi Client) GetType(id int64) (UniverseType, error) {
	var item UniverseType
	err := esi.get(fmt.Sprintf("/universe/types/%d/", id), &item)
	if err != nil {
		return UniverseType{}, err
	}

	return item, nil
}

// Gets information on an item group
func (esi Client) GetGroup(id int64) (GroupInfo, error) {
	var group GroupInfo
	err := esi.get(fmt.Sprintf("/universe/groups/%d/", id), &group)
	if err != nil {
		return GroupInfo{}, err
	}

	return group, nil
}

func (esi Client) GetSystems() ([]int64, error) {
	return esi.getIds("/universe/systems/")
}

func (esi Client) GetConstellations() ([]int64, error) {
	return esi.getIds("/universe/constellations/")
}

func (esi Client) GetRegions() ([]int64, error) {
	return esi.getIds("/universe/regions/")
}

func (esi Client) GetSystem(id int64) (SolarSystem, error) {
	var system SolarSystem
	err := esi.get(fmt.Sprintf("/universe/systems/%d/", id), &system)
	if err != nil {
		return SolarSystem{}, err
	}

	return system, nil
}

func (esi Client) GetConstellation(id int64) (Constellation, error) {
	var constellation Constellation
	err := esi.get(fmt.Sprintf("/universe/constellations/%d/", id), &constellation)
	if err != nil {
		return Constellation{}, err
	}

	return constellation, nil
}

func (esi Client) GetRegion(id int64) (Region, error) {
	var region Region
	err := esi.get(fmt.Sprintf("/universe/regions/%d/", id), &region)
	if err != nil {
		return Region{}, err
	}

	return region, nil
}

func (esi Client) GetStargate(id int64) (Stargate, error) {
	var gate Stargate
	err := esi.get(fmt.Sprintf("/universe/stargates/%d/", id), &gate)
	if err != nil {
		return Stargate{}, err
	}

	return gate, nil
}

func (esi Client) GetStation(id int64) (Station, error) {
	var station Station
	err := esi.get(fmt.Sprintf("/universe/stations/%d/", id), &station)
	if err != nil {
		return Station{}, err
	}

	return station, nil
}

func (esi Client) GetStar(id int64) (Star, error) {
	var star Star
	err := esi.get(fmt.Sprintf("/universe/stars/%d/", id), &star)
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
	err = esi.post("/universe/names/", buffer, &names)
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
