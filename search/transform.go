package search

import (
	"errors"
	"sort"
	"strings"

	"github.com/nfisher/swapi"
)

type EntityKey struct {
	EntityType string
	EntityId   string
}

type CorpusPair struct {
	Keys     []EntityKey
	Corpus   []string
	Universe *swapi.JsonUniverse
}

var ErrNotFound = errors.New("unable to retrieve item from search")

func (c CorpusPair) Retrieve(i int) (interface{}, error) {
	k := c.Keys[i]

	switch k.EntityType {
	case "film":
		v, ok := c.Universe.Films[k.EntityId]
		if !ok {
			return nil, ErrNotFound
		}
		return v, nil

	case "person":
		v, ok := c.Universe.People[k.EntityId]
		if !ok {
			return nil, ErrNotFound
		}
		return v, nil

	case "planet":
		v, ok := c.Universe.Planets[k.EntityId]
		if !ok {
			return nil, ErrNotFound
		}
		return v, nil

	case "specie":
		v, ok := c.Universe.Species[k.EntityId]
		if !ok {
			return nil, ErrNotFound
		}
		return v, nil

	case "starship":
		v, ok := c.Universe.Starships[k.EntityId]
		if !ok {
			return nil, ErrNotFound
		}
		return v, nil

	case "vehicle":
		v, ok := c.Universe.Vehicles[k.EntityId]
		if !ok {
			return nil, ErrNotFound
		}
		return v, nil
	}

	return nil, ErrNotFound
}

// NewCorpus creates an index based on the universe and returns a CorpusPair.
// TODO: This is a fucking mess!!
func NewCorpus(u *swapi.JsonUniverse) *CorpusPair {
	var pair = CorpusPair{
		Universe: u,
	}

	var keys []string
	for k := range u.Films {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	kind := "film"
	for _, k := range keys {
		c := u.Films[k]
		pair.Keys = append(pair.Keys, EntityKey{EntityType: kind, EntityId: c.Id})
		pair.Corpus = append(pair.Corpus, strings.ToLower(c.String()))
	}

	keys = make([]string, 0, len(u.People))
	for k := range u.People {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	kind = "person"
	for _, k := range keys {
		c := u.People[k]
		pair.Keys = append(pair.Keys, EntityKey{EntityType: kind, EntityId: c.Id})
		pair.Corpus = append(pair.Corpus, strings.ToLower(c.String()))
	}

	keys = make([]string, 0, len(u.Planets))
	for k := range u.Planets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	kind = "planet"
	for _, k := range keys {
		c := u.Planets[k]
		pair.Keys = append(pair.Keys, EntityKey{EntityType: kind, EntityId: c.Id})
		pair.Corpus = append(pair.Corpus, strings.ToLower(c.String()))
	}

	keys = make([]string, 0, len(u.Species))
	for k := range u.Species {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	kind = "specie"
	for _, k := range keys {
		c := u.Species[k]
		pair.Keys = append(pair.Keys, EntityKey{EntityType: kind, EntityId: c.Id})
		pair.Corpus = append(pair.Corpus, strings.ToLower(c.String()))
	}

	keys = make([]string, 0, len(u.Starships))
	for k := range u.Starships {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	kind = "starship"
	for _, k := range keys {
		c := u.Starships[k]
		pair.Keys = append(pair.Keys, EntityKey{EntityType: kind, EntityId: c.Id})
		pair.Corpus = append(pair.Corpus, strings.ToLower(c.String()))
	}

	keys = make([]string, 0, len(u.Vehicles))
	for k := range u.Vehicles {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	kind = "vehicle"
	for _, k := range keys {
		c := u.Vehicles[k]
		pair.Keys = append(pair.Keys, EntityKey{EntityType: kind, EntityId: c.Id})
		pair.Corpus = append(pair.Corpus, strings.ToLower(c.String()))
	}

	return &pair
}
