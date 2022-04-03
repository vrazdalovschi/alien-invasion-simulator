package domain

import (
	"fmt"
	"strings"
)

type Direction string

const (
	North Direction = "north"
	East  Direction = "east"
	South Direction = "south"
	West  Direction = "west"
)

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	}
	return ""
}

func (d Direction) String() string {
	return string(d)
}

func (d *Direction) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "north":
		*d = North
	case "east":
		*d = East
	case "south":
		*d = South
	case "west":
		*d = West
	default:
		return fmt.Errorf("invalid direction: %s", text)
	}
	return nil
}
