package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vrazdalovschi/alien-invasion-simulator/internal/domain"
)

func TestAlien_ChooseNextCity(t *testing.T) {
	t.Run("city without directions", func(t *testing.T) {
		city := &domain.City{
			Name:       "Foo",
			Directions: map[domain.Direction]*domain.City{},
		}
		alien := domain.Alien{
			City: city,
		}
		nextCity := alien.ChooseNextCity()
		assert.Equal(t, city, nextCity)
	})
	t.Run("city with 1 direction", func(t *testing.T) {
		fooCity := &domain.City{
			Name:       "Foo",
			Directions: map[domain.Direction]*domain.City{},
		}
		barCity := &domain.City{
			Name:       "Bar",
			Directions: map[domain.Direction]*domain.City{},
		}
		fooCity.Directions[domain.North] = barCity
		barCity.Directions[domain.South] = fooCity

		alien := domain.Alien{
			City: fooCity,
		}

		nextCity := alien.ChooseNextCity()
		assert.Equal(t, barCity, nextCity)
		assert.NotEqual(t, fooCity, nextCity)

		// move alien to barCity, it will require to choose next city (current city) again
		alien.MoveTo(barCity)

		nextCity = alien.ChooseNextCity()
		assert.Equal(t, fooCity, nextCity)
	})

	// probability of choosing a city is 1/4, it can be tested with a rand.Seed(1), but it's not necessary
	// this test iterates up to 1000 times and it should be enough to ensure that every city is "choosen" at least once
	t.Run("city with 4 directions", func(t *testing.T) {
		fooCity := &domain.City{Name: "Foo"}
		barCity := &domain.City{Name: "Bar"}
		bazCity := &domain.City{Name: "Baz"}
		quuxCity := &domain.City{Name: "Qu-ux"}
		beeCity := &domain.City{Name: "Bee"}

		fooCity.Directions = map[domain.Direction]*domain.City{
			domain.North: barCity,
			domain.West:  bazCity,
			domain.South: quuxCity,
			domain.East:  beeCity,
		}
		alien := domain.Alien{
			City: fooCity,
		}

		expectedCities := map[*domain.City]bool{
			barCity:  true,
			bazCity:  true,
			quuxCity: true,
			beeCity:  true,
		}

		for i := 0; i < 1000; i++ {
			nextCity := alien.ChooseNextCity()
			delete(expectedCities, nextCity)
			if len(expectedCities) == 0 {
				break
			}
		}

		assert.Equal(t, 0, len(expectedCities))
	})
}

func TestAlien_Move(t *testing.T) {
	var alien domain.Alien
	city := &domain.City{Name: "Foo"}
	alien.MoveTo(city)
	assert.Equal(t, city, alien.City)
	assert.Equal(t, 1, alien.Steps)
}
