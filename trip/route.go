package trip

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"googlemaps.github.io/maps"
)

type Configuration struct {
	GoogleAPIKey string
}

//apiKey                   "key", "", "API Key for using Google Maps API."
//clientID                 "client_id", "", "ClientID for Maps for Work API access."
//origin                   "origin", "", "The address or textual latitude/longitude value from which you wish to calculate directions."
//Destination              "Destination", "", "The address or textual latitude/longitude value from which you wish to calculate directions."
//Mode                     "Mode", "", "The travel Mode for this directions request."
//DepartureTime            "departure_time", "", "The depature time for transit Mode directions request."
//ArrivalTime              "arrival_time", "", "The arrival time for transit Mode directions request."
//Waypoints                "Waypoints", "", "The Waypoints for driving directions request, | separated."
//Language                 "Language", "", "Specifies the Language in which to return results."
//Region                   "Region", "", "Specifies the Region code, specified as a ccTLD (\"top-level domain\") two-character value."
//TransitMode              "transit_mode", "", "Specifies one or more preferred modes of transit, | separated. This parameter may only be specified for transit directions."
//TransitRoutingPreference "transit_routing_preference", "", "Specifies preferences for transit routes."
//iterations               "iterations", 1, "Number of times to make API request."
//TrafficModel             ("traffic_model", "", "Specifies traffic prediction model when request future directions. Valid values are optimistic, best_guess, and pessimistic. Optional."
type GoogleCustomRouteRequest struct {
	Origin                   string
	Destination              string
	Mode                     string
	DepartureTime            string
	ArrivalTime              string
	Waypoints                string
	WaypointsTime			 string
	Language                 string
	Region                   string
	TransitMode              string
	TransitRoutingPreference string
	TrafficModel             string
}

func analyzeWayppoints(request GoogleCustomRouteRequest) []maps.Route {
	r := &maps.DirectionsRequest{
		Origin:        request.Origin,
		Destination:   request.Destination,
		DepartureTime: request.DepartureTime,
		ArrivalTime:   request.ArrivalTime,
		Language:      request.Language,
		Region:        request.Region,
	}

	lookupMode(request.Mode, r)
	lookupMode(request.Mode, r)
	lookupTransitRoutingPreference(request.TransitRoutingPreference, r)
	lookupTrafficModel(request.TrafficModel, r)

	if request.Waypoints != "" {
		r.Waypoints = strings.Split(request.Waypoints, "|")
	}

	if request.TransitMode != "" {
		for _, t := range strings.Split(request.TransitMode, "|") {
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

	routes, _, err := client.Directions(context.Background(), r)
	check(err)

	// fmt.Printf("%# v", pretty.Formatter(routes))
	// fmt.Printf("%# v", pretty.Formatter(waypoints))
	return routes
}

func Route(request GoogleCustomRouteRequest) []maps.Route {
	configuration, _ := getSecrets()

	var client *maps.Client
	var err error
	if configuration.GoogleAPIKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(configuration.GoogleAPIKey), maps.WithRateLimit(2))
	} else {
		_, err = fmt.Fprintln(os.Stderr, "Please specify an API Key.")
		os.Exit(2)
	}
	check(err)

	r := &maps.DirectionsRequest{
		Origin:        request.Origin,
		Destination:   request.Destination,
		DepartureTime: request.DepartureTime,
		ArrivalTime:   request.ArrivalTime,
		Language:      request.Language,
		Region:        request.Region,
	}

	lookupMode(request.Mode, r)
	lookupMode(request.Mode, r)
	lookupTransitRoutingPreference(request.TransitRoutingPreference, r)
	lookupTrafficModel(request.TrafficModel, r)

	if request.Waypoints != "" {
		r.Waypoints = strings.Split(request.Waypoints, "|")
	}

	if request.TransitMode != "" {
		for _, t := range strings.Split(request.TransitMode, "|") {
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

	routes, _, err := client.Directions(context.Background(), r)
	check(err)

	// fmt.Printf("%# v", pretty.Formatter(routes))
	// fmt.Printf("%# v", pretty.Formatter(waypoints))
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
