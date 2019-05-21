package trip

import (
	"context"
	"log"
	"strings"

	"googlemaps.github.io/maps"
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
	Language                 string
	Region                   string
	TransitMode              string
	TransitRoutingPreference string
	TrafficModel             string
}

func Route(request GoogleCustomRouteRequest) []maps.Route {
	client := getGoogleClient()

	// Create directions request
	r := &maps.DirectionsRequest{
		Origin:        request.Origin,
		Destination:   request.Destination,
		DepartureTime: request.DepartureTime,
		ArrivalTime:   request.ArrivalTime,
		Language:      request.Language,
		Region:        request.Region,
		Waypoints:     request.Waypoints,
	}

	// fillGraph(request.Waypoints, request.WaypointsTime)

	// Set default mode as first selected
	if len(request.Mode) > 0 {
		lookupMode(request.Mode[0], r)
	} else {
		lookupMode("", r)
	}

	// set Google Directions request content (TransitRoutingPreference)
	lookupTransitRoutingPreference(request.TransitRoutingPreference, r)
	// set Google Directions request content (Mode)
	lookupTrafficModel(request.TrafficModel, r)
	// separate | transit mode and rest Google Directions request content (TransitMode)
	lookupTransitMode(request.TransitMode, r)

	// CALL DIRECTIONS API
	routes, _, err := client.Directions(context.Background(), r)
	check(err)

	//fmt.Printf("%# v", pretty.Formatter(routes))
	return routes

}

func lookupTransitMode(mode string, r *maps.DirectionsRequest) {
	if mode != "" {
		for _, t := range strings.Split(mode, "|") {
			switch t {
			case "bus":
				r.TransitMode = append(r.TransitMode, maps.TransitModeBus)
			case "subway":
				r.TransitMode = append(r.TransitMode, maps.TransitModeSubway)
			case "train":
				r.TransitMode = append(r.TransitMode, maps.TransitModeTrain)
			case "tram":
				r.TransitMode = append(r.TransitMode, maps.TransitModeTram)
			case "rail":
				r.TransitMode = append(r.TransitMode, maps.TransitModeRail)
			}
		}
	}
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

func lookupTransitRoutingPreference(transitRoutingPreference string, r *maps.DirectionsRequest) {
	switch transitRoutingPreference {
	case "fewer_transfers":
		r.TransitRoutingPreference = maps.TransitRoutingPreferenceFewerTransfers
	case "less_walking":
		r.TransitRoutingPreference = maps.TransitRoutingPreferenceLessWalking
	case "":
		// ignore
	default:
		log.Fatalf("Unknown transit routing preference %s", transitRoutingPreference)
	}
}

func lookupTrafficModel(trafficModel string, r *maps.DirectionsRequest) {
	switch trafficModel {
	case "optimistic":
		r.TrafficModel = maps.TrafficModelOptimistic
	case "best_guess":
		r.TrafficModel = maps.TrafficModelBestGuess
	case "pessimistic":
		r.TrafficModel = maps.TrafficModelPessimistic
	case "":
		// ignore
	default:
		log.Fatalf("Unknown traffic Mode %s", trafficModel)
	}
}
