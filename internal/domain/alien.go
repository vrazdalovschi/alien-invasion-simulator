package domain

import "math/rand"

type Alien struct {
	Name int

	// City is the city the alien is currently in
	City *City

	// Steps is the number of steps the alien has taken
	Steps int
}

// ChooseNextCity chooses a random city from the list of the current city directions
// if city doesn't have any directions, it returns the current city
func (a *Alien) ChooseNextCity() *City {
	// the city doesn't have direction, alien cannot move to any other city
	if len(a.City.Directions) == 0 {
		return a.City
	}

	var i int
	randomCityIndex := rand.Intn(len(a.City.Directions))

	for _, city := range a.City.Directions {
		if i == randomCityIndex {
			return city
		}
		i++
	}

	return a.City
}

// MoveTo update current city and increment steps
func (a *Alien) MoveTo(city *City) {
	a.City = city
	a.Steps++
}
