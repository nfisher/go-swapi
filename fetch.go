package swapi

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func open(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
}

var ErrEmpty = errors.New("Empty file contents")

// decode reads SW data from f which is a GOB cache if data is
// available.
func decode(f *os.File, e interface{}) error {
	stat, err := f.Stat()
	if err != nil {
		return err
	}

	if stat.Size() > 0 {
		if err = gob.NewDecoder(f).Decode(e); err != nil {
			return err
		}

		return nil
	}

	return ErrEmpty
}

// FetchRoot retrieves the root URL map from the SW API.
func FetchRoot(c *http.Client) (map[string]string, error) {
	f, err := open("root.gob")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m := make(map[string]string)
	err = decode(f, &m)
	if err != nil && err != ErrEmpty {
		return nil, err
	} else if err == nil {
		fmt.Printf("read %v entries from root.gob\n", len(m))
		return m, nil
	}

	// fetch from URL
	resp, err := c.Get("https://swapi.co/api/?format=json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}

	if gob.NewEncoder(f).Encode(m); err != nil {
		return nil, err
	}

	return m, nil
}

// API results

type Resultable interface {
	NextPage() string
	Reset()
	AppendTo(coll Collectable)
}

type Result struct {
	Count    int
	Next     string
	Previous string
}

func (r *Result) Reset() {
	r.Count = 0
	r.Next = ""
	r.Previous = ""
}

func (r *Result) NextPage() string {
	return r.Next
}

type PersonResult struct {
	Result
	Results People
}

func (r *PersonResult) Reset() {
	r.Result.Reset()
	r.Results = make(People, 0)
}

func (r *PersonResult) AppendTo(coll Collectable) {
	c := coll.(*People)
	c.Append(&r.Results)
}

type PlanetResult struct {
	Result
	Results Planets
}

func (r *PlanetResult) Reset() {
	r.Result.Reset()
	r.Results = make(Planets, 0)
}

func (r *PlanetResult) AppendTo(coll Collectable) {
	c := coll.(*Planets)
	c.Append(&r.Results)
}

type FilmResult struct {
	Result
	Results Films
}

func (r *FilmResult) Reset() {
	r.Result.Reset()
	r.Results = make(Films, 0)
}

func (r *FilmResult) AppendTo(coll Collectable) {
	c := coll.(*Films)
	c.Append(&r.Results)
}

type SpeciesResult struct {
	Result
	Results Species
}

func (r *SpeciesResult) Reset() {
	r.Result.Reset()
	r.Results = make(Species, 0)
}

func (r *SpeciesResult) AppendTo(coll Collectable) {
	c, ok := coll.(*Species)
	if !ok {
		return
	}
	c.Append(&r.Results)
}

type StarshipResult struct {
	Result
	Results Starships
}

func (r *StarshipResult) Reset() {
	r.Result.Reset()
	r.Results = make(Starships, 0)
}

func (r *StarshipResult) AppendTo(coll Collectable) {
	c := coll.(*Starships)
	c.Append(&r.Results)
}

type VehicleResult struct {
	Result
	Results Vehicles
}

func (r *VehicleResult) Reset() {
	r.Result.Reset()
	r.Results = make(Vehicles, 0)
}

func (r *VehicleResult) AppendTo(coll Collectable) {
	c := coll.(*Vehicles)
	c.Append(&r.Results)
}

func Fetch(c *http.Client, next, filename string, coll Collectable, result Resultable) error {
	if next == "" {
		return errors.New("no start URL provided")
	}

	f, err := open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// decode from the GOB file if possible
	err = decode(f, coll)
	if err != nil && err != ErrEmpty {
		return err
	} else if err == nil && coll.Len() > 0 {
		fmt.Printf("read %v entries from %s\n", coll.Len(), filename)
		return nil
	}

	// no entries found, download from the Star Wars API.
	for next != "" {
		fmt.Println(next)
		resp, err := c.Get(next)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
			return err
		}

		result.AppendTo(coll)

		next = result.NextPage()
		result.Reset()
	}

	// write the results out to the GOB file
	if err = gob.NewEncoder(f).Encode(coll); err != nil {
		return err
	}

	return nil
}
