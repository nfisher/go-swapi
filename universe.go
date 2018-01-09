package swapi

import (
	"encoding/json"
	"errors"
	"os"
)

type Universe struct {
	Films     Films
	People    People
	Planets   Planets
	Species   Species
	Starships Starships
	Vehicles  Vehicles
}

func LoadUniverse(filename string) (*JsonUniverse, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var ju JsonUniverse

	if err = json.NewDecoder(f).Decode(&ju); err != nil {
		return nil, err
	}

	return &ju, nil
}

type JsonUniverse struct {
	Films     map[string]Film
	People    map[string]Person
	Planets   map[string]Planet
	Species   map[string]Specie
	Starships map[string]Starship
	Vehicles  map[string]Vehicle
}

var FilmNotFound = errors.New("Film not found")
var PersonNotFound = errors.New("Person not found")
var PlanetNotFound = errors.New("Planet not found. Death star nearby?")
var SpecieNotFound = errors.New("Specie not found")
var StarshipNotFound = errors.New("Starship not found")
var VehicleNotFound = errors.New("Vehicle not found")

func (ju *JsonUniverse) Film(id string) (*Film, error) {
	f, ok := ju.Films[id]
	if !ok {
		return nil, FilmNotFound
	}

	return &f, nil
}

func (ju *JsonUniverse) Person(id string) (*Person, error) {
	f, ok := ju.People[id]
	if !ok {
		return nil, PersonNotFound
	}

	return &f, nil
}

func (ju *JsonUniverse) Planet(id string) (*Planet, error) {
	f, ok := ju.Planets[id]
	if !ok {
		return nil, SpecieNotFound
	}

	return &f, nil
}

func (ju *JsonUniverse) Specie(id string) (*Specie, error) {
	f, ok := ju.Species[id]
	if !ok {
		return nil, SpecieNotFound
	}

	return &f, nil
}

func (ju *JsonUniverse) Starship(id string) (*Starship, error) {
	f, ok := ju.Starships[id]
	if !ok {
		return nil, StarshipNotFound
	}

	return &f, nil
}

func (ju *JsonUniverse) Vehicle(id string) (*Vehicle, error) {
	f, ok := ju.Vehicles[id]
	if !ok {
		return nil, VehicleNotFound
	}

	return &f, nil
}

func Struct2map(u *Universe) *JsonUniverse {
	var ju JsonUniverse
	// should probably handle dup keys but can't be faffed
	ju.Films = make(map[string]Film)
	for _, v := range u.Films {
		id := string(v.Url)
		v.Id = id
		v.Url = "" // drop the URL in the JSON
		ju.Films[id] = v
	}

	ju.People = make(map[string]Person)
	for _, v := range u.People {
		id := string(v.Url)
		v.Id = id
		v.Url = ""
		ju.People[id] = v
	}

	ju.Planets = make(map[string]Planet)
	for _, v := range u.Planets {
		id := string(v.Url)
		v.Id = id
		v.Url = ""
		ju.Planets[id] = v
	}

	ju.Species = make(map[string]Specie)
	for _, v := range u.Species {
		id := string(v.Url)
		v.Id = id
		v.Url = ""
		ju.Species[id] = v
	}

	ju.Starships = make(map[string]Starship)
	for _, v := range u.Starships {
		id := string(v.Url)
		v.Id = id
		v.Url = ""
		ju.Starships[id] = v
	}

	ju.Vehicles = make(map[string]Vehicle)
	for _, v := range u.Vehicles {
		id := string(v.Url)
		v.Id = id
		v.Url = ""
		ju.Vehicles[id] = v
	}

	return &ju
}
