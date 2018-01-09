package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/nfisher/swapi"
)

func main() {
	client := &http.Client{}
	root, err := swapi.FetchRoot(client)
	if err != nil {
		fmt.Errorf("%v\n", err)
		return
	}

	var universe swapi.Universe

	// could make this parallel but feels a little douchy to hammer
	// their API. Could probably clean this up some too
	entities := []struct {
		basename string
		data     swapi.Collectable
		result   swapi.Resultable
	}{
		{"films", &universe.Films, &swapi.FilmResult{}},
		{"people", &universe.People, &swapi.PersonResult{}},
		{"planets", &universe.Planets, &swapi.PlanetResult{}},
		{"species", &universe.Species, &swapi.SpeciesResult{}},
		{"starships", &universe.Starships, &swapi.StarshipResult{}},
		{"vehicles", &universe.Vehicles, &swapi.VehicleResult{}},
	}

	for _, v := range entities {
		filename := v.basename + ".gob"
		err = swapi.Fetch(client, root[v.basename], filename, v.data, v.result)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	f, err := os.Create("swapi.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	m := swapi.Struct2map(&universe)

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err = enc.Encode(&m); err != nil {
		fmt.Println(err)
		return
	}
}
