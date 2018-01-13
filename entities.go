package swapi

import (
	"encoding/json"
	"fmt"
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
	Append(coll Collectable)
	Len() int
}

type People []Person

func (p *People) Append(coll Collectable) {
	res, ok := coll.(*People)
	if !ok {
		return
	}
	*p = append(*p, *res...)
}

func (p *People) Len() int {
	return len(*p)
}

type Planets []Planet

func (p *Planets) Append(coll Collectable) {
	res, ok := coll.(*Planets)
	if !ok {
		return
	}
	*p = append(*p, *res...)
}

func (p *Planets) Len() int {
	return len(*p)
}

type Films []Film

func (f *Films) Append(coll Collectable) {
	res, ok := coll.(*Films)
	if !ok {
		return
	}
	*f = append(*f, *res...)
}

func (f *Films) Len() int {
	return len(*f)
}

type Species []Specie

func (f *Species) Append(coll Collectable) {
	res, ok := coll.(*Species)
	if !ok {
		return
	}
	*f = append(*f, *res...)
}

func (f *Species) Len() int {
	return len(*f)
}

type Starships []Starship

func (f *Starships) Append(coll Collectable) {
	res, ok := coll.(*Starships)
	if !ok {
		return
	}
	*f = append(*f, *res...)
}

func (f *Starships) Len() int {
	return len(*f)
}

type Vehicles []Vehicle

func (f *Vehicles) Append(coll Collectable) {
	res, ok := coll.(*Vehicles)
	if !ok {
		return
	}
	*f = append(*f, *res...)
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
	Homeworld Url
	Films     Ids
	Species   Ids
	Vehicles  Ids
	Starships Ids
	RestFields
}

func (p Person) String() string {
	return fmt.Sprintf("%#v", p)
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

func (p Planet) String() string {
	return fmt.Sprintf("%#v", p)
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

func (f Film) String() string {
	return fmt.Sprintf("%v %v %v %v %v", f.Title, f.EpisodeID, f.Director, f.Producer, f.ReleaseDate)
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
	Homeworld       Url
	People          Ids
	Films           Ids
	RestFields
}

func (s Specie) String() string {
	return fmt.Sprintf("%#v", s)
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

func (s Starship) String() string {
	return fmt.Sprintf("%#v", s)
}

type Vehicle struct {
	Transport

	VehicleClass string `json:"vehicle_class"`
}

func (v Vehicle) String() string {
	return fmt.Sprintf("%#v", v)
}
