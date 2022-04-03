package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vrazdalovschi/alien-invasion-simulator/internal/domain"
)

func TestCity_String(t *testing.T) {
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

	actual := fooCity.String()
	assert.Equal(t, "Foo north=Bar west=Baz south=Qu-ux east=Bee", actual)
}
