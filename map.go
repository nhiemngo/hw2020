package main

import (
	"context"
	"fmt"
	"os"

	"googlemaps.github.io/maps"
)

var googleApiKey string = os.Getenv("GOOGLE_API_KEY")

func getGeocoding(address string) (latitude, longitude float64) {
	c, err := maps.NewClient(maps.WithAPIKey(googleApiKey))
	if err != nil {
		fmt.Println("fatal error: %s", err)
	}

	r := &maps.GeocodingRequest{
		Address: address,
	}

	location, err := c.Geocode(context.Background(), r)
	if err != nil {
		fmt.Println("fatal error: %s", err)
	}

	return location[0].Geometry.Location.Lat, location[0].Geometry.Location.Lng
}
