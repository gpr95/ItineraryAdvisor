package trip

import (
	"context"
	"googlemaps.github.io/maps"
	"log"
)


// Origin     				"The address or textual latitude/longitude value from which you wish to calculate directions."
// Destination              "The address or textual latitude/longitude value from which you wish to calculate directions."
// Mode                     "The travel Mode for this directions request."
// DepartureTime            "The depature time for transit Mode directions request."
// ArrivalTime              "The arrival time for transit Mode directions request."
// Waypoints                "The Waypoints for driving directions request, | separated."
// WaypointsTime            "The time planned to spend in certain waypoint, | separated."
// Language                 "Specifies the Language in which to return results."
// Region                   "Specifies the Region code, specified as a ccTLD (\"top-level domain\") two-character value."
// TransitMode              "Specifies one or more preferred modes of transit, | separated. This parameter may only be specified for transit directions."
// TransitRoutingPreference "Specifies preferences for transit routes."
// TrafficModel             "Specifies traffic prediction model when request future directions.
// 							Valid values are optimistic, best_guess, and pessimistic. Optional."
type GoogleCustomRouteRequest struct {
	Origin                   string
	Destination              string
	Mode                     []string
	DepartureTime            string
	ArrivalTime              string
	Waypoints                []string
	WaypointsTime            []string
}

func Route(request GoogleCustomRouteRequest) []maps.Route {
	client := getGoogleClient()

	// Create directions request
	r := &maps.DirectionsRequest{
		Origin:        request.Origin,
		Destination:   request.Destination,
		DepartureTime: request.DepartureTime,
		ArrivalTime:   request.ArrivalTime,
		Language:      "PL",
		Waypoints:     request.Waypoints,
	}

	// fillGraph(request.Waypoints, request.WaypointsTime)

	// Set default mode as first selected
	if len(request.Mode) > 0 {
		lookupMode(request.Mode[0], r)
	} else {
		lookupMode("", r)
	}

	// CALL DIRECTIONS API
	routes, _, err := client.Directions(context.Background(), r)
	check(err)

	//fmt.Printf("%# v", pretty.Formatter(routes))
	return routes

}


func lookupMode(mode string, r *maps.DirectionsRequest) {
	switch mode {
	case "driving":
		r.Mode = maps.TravelModeDriving
	case "walking":
		r.Mode = maps.TravelModeWalking
	case "bicycling":
		r.Mode = maps.TravelModeBicycling
	case "transit":
		r.Mode = maps.TravelModeTransit
	case "":
		// ignore
	default:
		log.Fatalf("Unknown Mode '%s'", mode)
	}
}
