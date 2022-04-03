package domain

import "strings"

type City struct {
	// Name of the city, any characters except space.
	Name string

	// Directions to the connected cities.
	Directions map[Direction]*City
}

// String format the city to the file format:
// <city> <direction>=<destinationCity>
// Example:
// 		Foo north=Baz
func (c *City) String() string {
	var buffer strings.Builder
	buffer.WriteString(c.Name)
	for direction, city := range c.Directions {
		buffer.WriteString(" ")
		buffer.WriteString(direction.String())
		buffer.WriteString("=")
		buffer.WriteString(city.Name)
	}
	return buffer.String()
}
