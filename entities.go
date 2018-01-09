package swapi

import (
	"encoding/json"
	"strings"
)

func url2id(url string) string {
	if len(url) < 4 {
		return url
	}
	last := len(url) - 1
	pos := strings.LastIndex(url[:last], "/")
	return url[pos+1 : last]
}

// Entity Collections

type Collectable interface {
	Len() int
}

type People []Person

func (p *People) Append(result interface{}) {
	res := result.(*PersonResult)
	*p = append(*p, res.Results...)
}

func (p *People) Len() int {
	return len(*p)
}

type Planets []Planet

func (p *Planets) Append(result interface{}) {
	res := result.(*PlanetResult)
	*p = append(*p, res.Results...)
}

func (p *Planets) Len() int {
	return len(*p)
}

type Films []Film

func (f *Films) Append(result interface{}) {
	res := result.(*FilmResult)
	*f = append(*f, res.Results...)
}

func (f *Films) Len() int {
	return len(*f)
}

type Species []Specie

func (f *Species) Append(result interface{}) {
	res := result.(*SpeciesResult)
	*f = append(*f, res.Results...)
}

func (f *Species) Len() int {
	return len(*f)
}

type Starships []Starship

func (f *Starships) Append(result interface{}) {
	res := result.(*StarshipResult)
	*f = append(*f, res.Results...)
}

func (f *Starships) Len() int {
	return len(*f)
}

type Vehicles []Vehicle

func (f *Vehicles) Append(result interface{}) {
	res := result.(*VehicleResult)
	*f = append(*f, res.Results...)
}

func (f *Vehicles) Len() int {
	return len(*f)
}

// Entities
type Ids []string

func (fi *Ids) UnmarshalJSON(b []byte) error {
	var urls []string
	err := json.Unmarshal(b, &urls)
	if err != nil {
		return err
	}

	for i, v := range urls {
		urls[i] = url2id(v)
	}

	*fi = urls

	return nil
}

type Url string

func (ptr *Url) UnmarshalJSON(b []byte) error {
	var u string
	err := json.Unmarshal(b, &u)
	if err != nil {
		return err
	}

	*ptr = Url(url2id(u))

	return nil
}

type RestFields struct {
	Id      string
	Created string
	Edited  string
	Url     Url `json:",omitempty"`
}

type Person struct {
	Name      string
	Height    string
	Mass      string
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string
	Homeworld string
	Films     Ids
	Species   Ids
	Vehicles  Ids
	Starships Ids
	RestFields
}

type Planet struct {
	Name           string
	Diameter       string
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Gravity        string
	Population     string
	Climate        string
	Terrain        string
	SurfaceWater   string `json:"surface_water"`
	Residents      Ids
	Films          Ids
	RestFields
}

type Film struct {
	Title        string
	EpisodeID    int    `json:"episode_id"`
	OpeningCrawl string `json:"opening_crawl"`
	Director     string
	Producer     string
	ReleaseDate  string `json:"release_date"`
	Species      Ids
	Starships    Ids
	Vehicles     Ids
	Characters   Ids
	Planets      Ids
	RestFields
}

type Specie struct {
	Name            string
	Classification  string
	Designation     string
	AverageHeight   string `json:"average_height"`
	AverageLifespan string `json:"average_lifespan"`
	EyeColors       string `json:"eye_colors"`
	HairColors      string `json:"hair_colors"`
	SkinColors      string `json:"skin_colors"`
	Language        string
	Homeworld       string
	People          Ids
	Films           Ids
	RestFields
}

type Transport struct {
	CargoCapacity        string `json:"cargo_capacity"`
	Consumables          string
	CostInCredits        string `json:"cost_in_credits"`
	Crew                 string
	Films                Ids
	Length               string
	Manufacturer         string
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
	Model                string
	Name                 string
	Passengers           string
	Pilots               Ids

	RestFields
}

type Starship struct {
	Transport

	Mglt             string
	HyperdriveRating string `json:"hyperdrive_rating"`
	StarshipClass    string `json:"starship_class"`
}

type Vehicle struct {
	Transport

	VehicleClass string `json:"vehicle_class"`
}
