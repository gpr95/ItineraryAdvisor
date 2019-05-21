package trip

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
	"time"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

// FrontendResponse is struct that frontend will understand
type FrontendResponse struct {
	Route            []Step
	Distance         string
	Duration         time.Duration
	ArrivalTime      time.Time
	DepartureTime    time.Time
	OverviewPolyline maps.Polyline
}

type Step struct {
	Location    maps.LatLng
	Instruction string
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
		meters := int(0)
		for _, leg := range route.Legs {
			fmt.Println("Leg distance -> " + leg.Distance.HumanReadable)
			meters = leg.Distance.Meters + meters
			fmt.Println("Total distance -> " + strconv.Itoa(meters) + " m")

			fmt.Println("Leg duration -> " + leg.Duration.String())
			output.Duration = output.Duration + leg.Duration
			fmt.Println("Total duration -> " + output.Duration.String())
			output.ArrivalTime = leg.ArrivalTime
			output.DepartureTime = leg.DepartureTime
			for _, step := range leg.Steps {
				newStep := Step{Location: step.StartLocation, Instruction: step.HTMLInstructions}
				output.Route = append(output.Route, newStep)
			}
		}
		output.Distance = strconv.Itoa(meters) + " m"
		if len(output.Distance) > 3 {
			strMeters := output.Distance[len(output.Distance)-5 : len(output.Distance)-1]
			strKm := output.Distance[0 : len(output.Distance)-5]

			output.Distance = strKm + " km " + strMeters + " m"
		}
	}
	return output
}

// ParseFrontendRequest parses frontend request, and returns GoogleCustomRouteRequest
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
		case "lookup-mode":
			var lookupModeList []lookupModeStruct
			_ = json.Unmarshal([]byte(value[0]), &lookupModeList)
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
			var waypointList []waypoint
			_ = json.Unmarshal([]byte(value[0]), &waypointList)
			for _, value := range waypointList {
				googleRequest.Waypoints = append(googleRequest.Waypoints, value.Name)
				googleRequest.WaypointsTime = append(googleRequest.WaypointsTime, value.Time)
			}
			googleRequest.Destination = googleRequest.Waypoints[len(googleRequest.Waypoints)-1]
		case "bounds":
			var directionMap map[string]float64
			_ = json.Unmarshal([]byte(value[0]), &directionMap)
			bounds := maps.LatLngBounds{
				NorthEast: maps.LatLng{
					Lat: directionMap["north"],
					Lng: directionMap["east"],
				},
				SouthWest: maps.LatLng{
					Lat: directionMap["south"],
					Lng: directionMap["west"],
				},
			}
			latitude :=
				strconv.FormatFloat((directionMap["north"]+directionMap["south"])/2.0, 'f', 6, 64) +
					","+
					strconv.FormatFloat((directionMap["west"]+directionMap["east"])/2.0, 'f', 6, 64)

			googleRequest := GoogleCustomNearbySearchRequest{
				 Location: latitude,
				 RankBy: "distance",
				 PlaceTypes: "museum",
			}
			NearbySearch(googleRequest)
			fmt.Printf("%# v", pretty.Formatter(bounds))
		}
	}
	return googleRequest
}

func ParseFetchPlacesRequest(c *gin.Context) GoogleCustomNearbySearchRequest {
	q := c.Request.URL.Query()
	north, _ := strconv.ParseFloat(q["north"][0],64)
	south, _ := strconv.ParseFloat(q["south"][0],64)
	west, _ := strconv.ParseFloat(q["west"][0],64)
	east, _ := strconv.ParseFloat(q["east"][0],64)
	println(north)
	latitude :=
		strconv.FormatFloat((north+south)/2.0, 'f', 6, 64) +
			","+
			strconv.FormatFloat((west+east)/2.0, 'f', 6, 64)
	println(latitude)
	return GoogleCustomNearbySearchRequest{
		Location: latitude,
		RankBy: "distance",
		PlaceTypes: "museum",
	}
}
