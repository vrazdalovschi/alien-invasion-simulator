package simulation

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/vrazdalovschi/alien-invasion-simulator/internal/domain"
)

type WorldConfiguration struct {
	// Aliens is the number of aliens to spawn
	Aliens int

	// MaxIterations is the maximum number of iterations to run the simulation
	MaxIterations int

	// Verbose is a flag to indicate printing every step of the simulation
	Verbose bool
}

type World struct {
	// Cities represent all cities in the map.
	Cities map[string]*domain.City

	// Aliens represent all aliens in the map.
	Aliens AliensSet

	// CityAliens tracks all aliens in the city.
	CityAliens map[*domain.City]AliensSet

	Configuration WorldConfiguration
}

type AliensSet map[*domain.Alien]bool

func (a AliensSet) String() string {
	identifiers := make([]string, 0, len(a))
	for alien := range a {
		identifiers = append(identifiers, strconv.Itoa(alien.Name))
	}
	return strings.Join(identifiers, ", ")
}

// NewWorld creates a new world and initialise maps.
func NewWorld(cfg WorldConfiguration) World {
	return World{
		Cities:        make(map[string]*domain.City),
		Aliens:        make(map[*domain.Alien]bool, cfg.Aliens),
		CityAliens:    make(map[*domain.City]AliensSet),
		Configuration: cfg,
	}
}

func (m *World) Simulate() error {
	if len(m.Cities) == 0 {
		return fmt.Errorf("no cities in the world")
	}
	if m.Configuration.Aliens == 0 {
		return fmt.Errorf("no aliens in the world")
	}

	m.placeAliens()

	for step := 0; step < m.Configuration.MaxIterations; step++ {
		if m.Configuration.Verbose {
			fmt.Println("Step:", step)
		}

		m.moveAliens()

		m.fightAliens(step)

		if len(m.Aliens) == 0 {
			fmt.Println("Step:", step, "All aliens are dead")
			break
		} else if len(m.Aliens) == 1 {
			fmt.Println("Step:", step, "Last one alien is alive, stop simulation")
			break
		}
	}

	fmt.Println("Simulation finished.")

	if m.Configuration.Verbose {
		fmt.Println("Statistics:")
		fmt.Println(" ", "* Aliens:", m.Aliens)
		fmt.Println(" ", "* World:")
		for _, city := range m.Cities {
			fmt.Println("   ", " *", city.String(), "Aliens:", m.CityAliens[city].String())
		}
	}

	return nil
}

// placeAliens create aliens and place them on the map
// TODO: refactor this method for custom AlienPlacement, that can be used for test seeding
func (m *World) placeAliens() {
	if m.Configuration.Verbose {
		fmt.Println("Placing aliens:")
	}
	for i := 0; i < m.Configuration.Aliens; i++ {
		city := m.getRandomCity()
		alien := &domain.Alien{Name: i, City: city}
		m.Aliens[alien] = true
		if _, ok := m.CityAliens[city]; !ok {
			m.CityAliens[city] = make(AliensSet)
		}
		m.CityAliens[city][alien] = true

		if m.Configuration.Verbose {
			fmt.Println("  * Alien", alien.Name, "placed in", city.Name)
		}
	}
}

// moveAliens moves all aliens to a random city (if city has any road)
func (m *World) moveAliens() {
	for alien := range m.Aliens {
		previousCity := alien.City
		nextCity := alien.ChooseNextCity()
		alien.MoveTo(nextCity)

		if previousCity != nextCity {
			if m.Configuration.Verbose {
				fmt.Println(" ", "* Alien", alien.Name, "moved from", previousCity.Name, "to", nextCity.Name)
			}
			// remove alien from the previous city
			delete(m.CityAliens[previousCity], alien)
			// add alien to new city
			if _, ok := m.CityAliens[nextCity]; !ok {
				m.CityAliens[nextCity] = make(AliensSet)
			}
			m.CityAliens[nextCity][alien] = true
		}
	}
}

// fightAliens destroy the city when 2 or more aliens are in the same city.
func (m *World) fightAliens(step int) {
	for city, cityAliens := range m.CityAliens {
		if len(cityAliens) > 1 {
			fmt.Println(fmt.Sprintf("Step: %d, City: %s has been destroyed by cityAliens: %s", step, city.Name, cityAliens.String()))
			for cityAlien := range cityAliens {
				delete(m.Aliens, cityAlien)
			}
			delete(m.CityAliens, city)

			// Remove roads from the city
			for direction, directionCity := range city.Directions {
				delete(directionCity.Directions, direction.Opposite())
			}
		}
	}
}

func (m *World) getRandomCity() *domain.City {
	randCityN := rand.Intn(len(m.Cities))
	j := 0
	for _, city := range m.Cities {
		if j == randCityN {
			return city
		}
		j++
	}
	return nil
}

func (m *World) ReadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := m.ReadLine(line); err != nil {
			return fmt.Errorf("error reading line %s, err: %w", line, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// ReadLine reads a city line, update the current connections between cities and tracks all cities.
// Line format: <city> <Direction>=<connectedCity>
// where:
//	* <city> is the name of the city
//	* <Direction> is the direction of the connection, see Direction
//  * <connectedCity> is the name of the connected city
// Example:
//	Foo north=Bar west=Baz south=Qu-ux
func (m *World) ReadLine(line string) error {
	parts := strings.Split(line, " ")
	cityName := parts[0]

	city, ok := m.Cities[cityName]
	if !ok {
		city = &domain.City{
			Name: cityName,
		}
		m.Cities[cityName] = city
	}

	if city.Directions == nil {
		city.Directions = make(map[domain.Direction]*domain.City)
	}

	for _, pair := range parts[1:] {
		directionParts := strings.Split(pair, "=")
		if len(directionParts) != 2 {
			return fmt.Errorf("invalid city=direction pair format: %s", pair)
		}
		directionName := directionParts[0]
		var direction domain.Direction
		if err := direction.UnmarshalText([]byte(directionName)); err != nil {
			return fmt.Errorf("failed to UnmarshalText direction: %w", err)
		}

		directionCityName := directionParts[1]
		if directionCityName == "" {
			return fmt.Errorf("empty direction city name for pair: %s", pair)
		}

		directionCity, ok := m.Cities[directionCityName]
		if !ok {
			directionCity = &domain.City{
				Name:       directionCityName,
				Directions: map[domain.Direction]*domain.City{},
			}
			m.Cities[directionCityName] = directionCity
		}

		if directionCityOpposite, ok := directionCity.Directions[direction.Opposite()]; ok && directionCityOpposite != city {
			return fmt.Errorf("city %s already has a connection to %s in direction %s", city.Name, directionCityOpposite.Name, direction)
		}

		directionCity.Directions[direction.Opposite()] = city
		city.Directions[direction] = directionCity
	}

	return nil
}
