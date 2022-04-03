package domain_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vrazdalovschi/alien-invasion-simulator/internal/domain"
)

func TestDirection_Opposite(t *testing.T) {
	testCases := []struct {
		In       domain.Direction
		Expected domain.Direction
	}{
		{domain.North, domain.South},
		{domain.South, domain.North},
		{domain.East, domain.West},
		{domain.West, domain.East},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s_%s", testCase.In, testCase.Expected), func(t *testing.T) {
			opposite := testCase.In.Opposite()
			assert.Equal(t, testCase.Expected, opposite)
		})
	}
}

func TestDirection_UnmarshalText(t *testing.T) {
	testCases := []struct {
		Name     string
		In       string
		Expected domain.Direction
		HasError bool
	}{
		{"lower case north", "north", domain.North, false},
		{"upper case north", "NORTH", domain.North, false},
		{"lower case south", "south", domain.South, false},
		{"mix case south", "souTh", domain.South, false},
		{"lower case east", "east", domain.East, false},
		{"lower case west", "west", domain.West, false},
		{"mix case west", "WEst", domain.West, false},
		{"unknown direction, double letter", "souuth", "", true},
		{"unknown direction", "northwest", "", true},
		{"empty direction", "", "", true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			var direction domain.Direction
			err := direction.UnmarshalText([]byte(testCase.In))
			if testCase.HasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.Expected, direction)
			}
		})
	}
}
