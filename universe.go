package swapi

type Universe struct {
	Films     Films
	People    People
	Planets   Planets
	Species   Species
	Starships Starships
	Vehicles  Vehicles
}

type JsonUniverse struct {
	Films     map[string]Film
	People    map[string]Person
	Planets   map[string]Planet
	Species   map[string]Specie
	Starships map[string]Starship
	Vehicles  map[string]Vehicle
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
