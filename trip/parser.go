package trip

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"googlemaps.github.io/maps"
)

// FrontendResponse is struct that frontend will understand
type FrontendResponse struct {
	Route            []maps.LatLng
	Distance         string
	Duration         time.Duration
	ArrivalTime      time.Time
	DepartureTime    time.Time
	OverviewPolyline maps.Polyline
}

type waypoint struct {
	Name string
	Time string
}

type lookupModeStruct struct {
	Name string
	Used bool
}

// GetCoordinatesAndInfoFromRoute creates response for frontend with route and some info
func GetCoordinatesAndInfoFromRoute(routes []maps.Route) FrontendResponse {
	var output FrontendResponse
	for _, route := range routes {
		output.OverviewPolyline = route.OverviewPolyline
		for _, leg := range route.Legs {
			output.Distance = leg.Distance.HumanReadable
			fmt.Println("Leg duration -> " + leg.Duration.String())
			output.Duration = output.Duration + leg.Duration
			fmt.Println("Total duration -> " + output.Duration.String())
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

// ParseFrontendRequest parses fronend request, and returns GoogleCustomRouteRequest
func ParseFrontendRequest(clientRequest url.Values) GoogleCustomRouteRequest {

	googleRequest := GoogleCustomRouteRequest{
		Origin:                   "",
		Destination:              "",
		Mode:                     []string{},
		DepartureTime:            "",
		ArrivalTime:              "",
		Waypoints:                []string{},
		WaypointsTime:            []string{},
		Language:                 "PL",
		Region:                   "",
		TransitMode:              "",
		TransitRoutingPreference: "",
		TrafficModel:             "",
	}

	for key, value := range clientRequest {
		fmt.Println(key, value)
		if value[0] == "undefined" || value[0] == "null" {
			continue
		}

		switch key {
		case "origin":
			googleRequest.Origin = value[0]
		case "destination":
			googleRequest.Destination = value[0]
		case "lookup-mode":
			lookupModeList := []lookupModeStruct{}
			json.Unmarshal([]byte(value[0]), &lookupModeList)
			for _, value := range lookupModeList {
				if value.Used {
					googleRequest.Mode = append(googleRequest.Mode, value.Name)
				}
			}
		case "departure":
			googleRequest.DepartureTime = value[0]
		case "arrival":
			googleRequest.ArrivalTime = value[0]
		case "waypoints":
			waypointList := []waypoint{}
			json.Unmarshal([]byte(value[0]), &waypointList)
			for _, value := range waypointList {
				googleRequest.Waypoints = append(googleRequest.Waypoints, value.Name)
				googleRequest.WaypointsTime = append(googleRequest.WaypointsTime, value.Time)
			}
		}
	}
	return googleRequest
}
