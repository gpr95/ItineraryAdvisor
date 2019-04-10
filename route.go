package main

import (
	"context"
	"fmt"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
	"log"
	"os"
	"strings"
)

type Configuration struct {
	GoogleAPIKey string
}

//apiKey                   "key", "", "API Key for using Google Maps API."
//clientID                 "client_id", "", "ClientID for Maps for Work API access."
//origin                   "origin", "", "The address or textual latitude/longitude value from which you wish to calculate directions."
//destination              "destination", "", "The address or textual latitude/longitude value from which you wish to calculate directions."
//mode                     "mode", "", "The travel mode for this directions request."
//departureTime            "departure_time", "", "The depature time for transit mode directions request."
//arrivalTime              "arrival_time", "", "The arrival time for transit mode directions request."
//waypoints                "waypoints", "", "The waypoints for driving directions request, | separated."
//language                 "language", "", "Specifies the language in which to return results."
//region                   "region", "", "Specifies the region code, specified as a ccTLD (\"top-level domain\") two-character value."
//transitMode              "transit_mode", "", "Specifies one or more preferred modes of transit, | separated. This parameter may only be specified for transit directions."
//transitRoutingPreference "transit_routing_preference", "", "Specifies preferences for transit routes."
//iterations               "iterations", 1, "Number of times to make API request."
//trafficModel             ("traffic_model", "", "Specifies traffic prediction model when request future directions. Valid values are optimistic, best_guess, and pessimistic. Optional."
type GoogleCustomRouteRequest struct {
	origin                   string
	destination              string
	mode                     string
	departureTime            string
	arrivalTime              string
	waypoints                string
	language                 string
	region                   string
	transitMode              string
	transitRoutingPreference string
	trafficModel             string
}

func route(request GoogleCustomRouteRequest) {
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
		Origin:        request.origin,
		Destination:   request.destination,
		DepartureTime: request.departureTime,
		ArrivalTime:   request.arrivalTime,
		Language:      request.language,
		Region:        request.region,
	}

	lookupMode(request.mode, r)
	lookupMode(request.mode, r)
	lookupTransitRoutingPreference(request.transitRoutingPreference, r)
	lookupTrafficModel(request.trafficModel, r)

	if request.waypoints != "" {
		r.Waypoints = strings.Split(request.waypoints, "|")
	}

	if request.transitMode != "" {
		for _, t := range strings.Split(request.transitMode, "|") {
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

	routes, waypoints, err := client.Directions(context.Background(), r)
	check(err)

	fmt.Printf("%# v", pretty.Formatter(waypoints))
	fmt.Printf("%# v", pretty.Formatter(routes))
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
		log.Fatalf("Unknown mode '%s'", mode)
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
		log.Fatalf("Unknown traffic mode %s", trafficModel)
	}
}
