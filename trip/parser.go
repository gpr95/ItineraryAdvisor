package trip

import (
	"googlemaps.github.io/maps"
)

type Waypoint struct {
	Lat string
	Lng string
}

func GetCoordinatesFromRoute(routes []maps.Route) []maps.LatLng {
	var output []maps.LatLng
	for _, route := range routes {
		for _, leg := range route.Legs {
			for _, step := range leg.Steps {
				output = append(output, step.StartLocation)
				output = append(output, step.EndLocation)
			}
		}
	}
	return output
}
