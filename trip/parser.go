package trip

import (
	"time"

	"googlemaps.github.io/maps"
)

// FrontendResponse ius struct that frontend will understand
type FrontendResponse struct {
	Route         []maps.LatLng
	Distance      string
	Duration      time.Duration
	ArrivalTime   time.Time
	DepartureTime time.Time
}

// GetCoordinatesAndInfoFromRoute creates response for frontend with route and some info
func GetCoordinatesAndInfoFromRoute(routes []maps.Route) FrontendResponse {
	var output FrontendResponse
	for _, route := range routes {
		for _, leg := range route.Legs {
			output.Distance = leg.Distance.HumanReadable
			output.Duration = leg.Duration
			output.ArrivalTime = leg.ArrivalTime
			output.DepartureTime = leg.DepartureTime
			for _, step := range leg.Steps {
				output.Route = append(output.Route, step.StartLocation)
				output.Route = append(output.Route, step.EndLocation)
			}
		}
	}
	return output
}
