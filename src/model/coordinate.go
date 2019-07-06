package model

import (
	"strconv"
)

type Coordinates []string

// TODO Can be improved to return more specific errors
func (c Coordinates) Valid() bool {
	if len(c) != 2 {
		return false
	}

	latitude, err := strconv.ParseFloat(c.Latitude(), 32)

	if err != nil || latitude < -90 || latitude > 90 {
		return false
	}

	longitude, err := strconv.ParseFloat(c.Longitude(), 32)

	if err != nil || longitude < -180 || longitude > 180 {
		return false
	}

	return true
}

func (c Coordinates) Latitude() string {
	return c[0]
}

func (c Coordinates) Longitude() string {
	return c[1]
}
