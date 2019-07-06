package service

import (
	"context"
	"errors"

	"github.com/vayan/sisistay/src/model"
	"googlemaps.github.io/maps"
)

type RouteFetcher interface {
	GetDistance(coordinates model.Coordinates, to model.Coordinates) (int, error)
}

type GoogleRouteFetcher struct {
	APIKey string
}

func (rf GoogleRouteFetcher) GetDistance(from model.Coordinates, to model.Coordinates) (int, error) {
	client, err := maps.NewClient(maps.WithAPIKey(rf.APIKey))

	if err != nil {
		return 0, errors.New("could not create Google Map client: " + err.Error())
	}
	r := &maps.DirectionsRequest{
		Origin:      from.ToString(),
		Destination: to.ToString(),
		Mode:        maps.TravelModeDriving, // Not good for the planet but assuming cars for now ;)
	}
	route, _, err := client.Directions(context.Background(), r)
	if err != nil {
		return 0, errors.New("could not get direction: " + err.Error())
	}

	if len(route) < 1 || len(route[0].Legs) < 1 {
		return 0, errors.New("could not find a route")
	}

	return route[0].Legs[0].Distance.Meters, nil
}
