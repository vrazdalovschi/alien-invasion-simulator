package simulation_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vrazdalovschi/alien-invasion-simulator/internal/domain"
	"github.com/vrazdalovschi/alien-invasion-simulator/internal/simulation"
)

func TestWorld_ReadLine(t *testing.T) {
	fooCity := &domain.City{Name: "Foo"}
	barCity := &domain.City{Name: "Bar"}
	bazCity := &domain.City{Name: "Baz"}
	quuxCity := &domain.City{Name: "Qu-ux"}
	beeCity := &domain.City{Name: "Bee"}

	fooCity.Directions = map[domain.Direction]*domain.City{
		domain.North: barCity,
		domain.West:  bazCity,
		domain.South: quuxCity,
	}
	barCity.Directions = map[domain.Direction]*domain.City{
		domain.South: fooCity,
		domain.West:  beeCity,
	}
	bazCity.Directions = map[domain.Direction]*domain.City{
		domain.East: fooCity,
	}
	quuxCity.Directions = map[domain.Direction]*domain.City{
		domain.North: fooCity,
	}
	beeCity.Directions = map[domain.Direction]*domain.City{
		domain.East: barCity,
	}

	expectedWorld := simulation.NewWorld(simulation.WorldConfiguration{})
	expectedWorld.Cities = map[string]*domain.City{
		"Foo":   fooCity,
		"Bar":   barCity,
		"Baz":   bazCity,
		"Qu-ux": quuxCity,
		"Bee":   beeCity,
	}

	actualWorld := simulation.NewWorld(simulation.WorldConfiguration{})
	lines := []string{
		"Foo north=Bar west=Baz south=Qu-ux",
		"Bar south=Foo west=Bee",
		"Baz east=Foo",
	}
	for _, line := range lines {
		err := actualWorld.ReadLine(line)
		assert.NoError(t, err)
	}

	assert.Equal(t, expectedWorld, actualWorld)
}

func TestWorld_ReadLine_Errors(t *testing.T) {
	testCases := []struct {
		Name string
		Line string
	}{
		{Name: "name with space", Line: "Fo o north=Bar west=Baz south=Qu-ux"},
		{Name: "direction with empty city", Line: "Foo north=Bar west= south=Qu-ux"},
		{Name: "unknown direction", Line: "Foo north=Bar hello=Qa south=Qu-ux"},
		{Name: "pair without equal sign", Line: "Foo north=Bar west=Baz southQu-ux"},
	}
	world := simulation.NewWorld(simulation.WorldConfiguration{})
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := world.ReadLine(testCase.Line)
			assert.Error(t, err)
		})
	}
}

// Cities are linked 1<>1 to each other via a direction.
// Validate that 2 cities cannot use the same direction to link with a city.
func TestWorld_ReadLine_WrongFlow(t *testing.T) {
	world := simulation.NewWorld(simulation.WorldConfiguration{})

	err := world.ReadLine("Foo")
	assert.NoError(t, err)

	err = world.ReadLine("Baz north=Foo")
	assert.NoError(t, err)

	err = world.ReadLine("Bar north=Foo")
	assert.Error(t, err)
}

func TestWorld_Simulate(t *testing.T) {
	t.Run("simple test, 1 city, kill them in the first step", func(t *testing.T) {
		world := simulation.NewWorld(simulation.WorldConfiguration{Aliens: 2, MaxIterations: 1})
		err := world.ReadLine("Foo")
		require.NoError(t, err)

		err = world.Simulate()
		require.NoError(t, err)

		assert.Equal(t, 0, len(world.Aliens))
	})
	t.Run("stop simulation when remain 1 alien", func(t *testing.T) {
		world := simulation.NewWorld(simulation.WorldConfiguration{Aliens: 1, MaxIterations: 10})
		err := world.ReadLine("Foo north=Bar west=Baz south=Qu-ux")
		require.NoError(t, err)

		err = world.Simulate()
		require.NoError(t, err)

		assert.Equal(t, 1, len(world.Aliens))
		for alien := range world.Aliens {
			assert.Equal(t, 1, alien.Steps)
		}
	})

}
