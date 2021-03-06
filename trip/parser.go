package trip

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

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
	Mode        string
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
func GetCoordinatesAndInfoFromRoute(routes []maps.Route, mode []string) FrontendResponse {
	var output FrontendResponse
	for _, route := range routes {
		output.OverviewPolyline = route.OverviewPolyline
		meters := int(0)
		for _, leg := range route.Legs {
			// fmt.Println("Leg distance -> " + leg.Distance.HumanReadable)
			meters = leg.Distance.Meters + meters
			// fmt.Println("Total distance -> " + strconv.Itoa(meters) + " m")

			// fmt.Println("Leg duration -> " + leg.Duration.String())
			output.Duration = output.Duration + leg.Duration
			// fmt.Println("Total duration -> " + output.Duration.String())
			output.ArrivalTime = leg.ArrivalTime
			output.DepartureTime = leg.DepartureTime
			for _, step := range leg.Steps {
				newStep := Step{Location: step.StartLocation, Instruction: step.HTMLInstructions, Mode: mode[0]}
				output.Route = append(output.Route, newStep)
			}
		}
		output.Distance = strconv.Itoa(meters) + " m"
		// if len(output.Distance) > 3 {
		// 	strMeters := output.Distance[len(output.Distance)-5 : len(output.Distance)-1]
		// 	strKm := output.Distance[0 : len(output.Distance)-5]

		// 	output.Distance = strKm + " km " + strMeters + " m"
		// }
	}
	return output
}

// ParseFrontendRequest parses frontend request, and returns GoogleCustomRouteRequest
func ParseFrontendRequest(clientRequest url.Values) GoogleCustomRouteRequest {

	googleRequest := GoogleCustomRouteRequest{
		Origin:        "",
		Destination:   "",
		Mode:          []string{},
		DepartureTime: "",
		ArrivalTime:   "",
		Waypoints:     []string{},
		WaypointsTime: []string{},
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
			if len(googleRequest.Waypoints) > 0 {
				googleRequest.Destination = googleRequest.Waypoints[len(googleRequest.Waypoints)-1]
			}
		}
	}
	return googleRequest
}

func ParseDateToCustomTimeString(date string) string {
	if date == "" {
		return "00:00"
	}
	removedDate := strings.Split(date, " ")[4]
	removedSeconds := strings.Split(removedDate, ":")[0] + ":" + strings.Split(removedDate, ":")[1]

	return removedSeconds
}

func ParseItineraryToGoogleRequests(placesList map[Place]TransportStatistics) []GoogleCustomRouteRequest {

	googleRequests := make([]GoogleCustomRouteRequest, 0)

	places := make([]Place, 0)
	for key := range placesList {
		places = append(places, key)
	}

	for index := 1; index < len(places); index++ {

		transitMode := placesList[places[index]].TransportType

		newGoogleRequest := GoogleCustomRouteRequest{
			Origin:        places[index-1].Name,
			Destination:   places[index].Name,
			Mode:          []string{transitMode},
			DepartureTime: "",
			ArrivalTime:   "",
			Waypoints:     []string{},
			WaypointsTime: []string{},
		}

		googleRequests = append(googleRequests, newGoogleRequest)
	}

	return googleRequests
}

func AppendGoogleResponse(base FrontendResponse, route []maps.Route, mode []string, request GoogleCustomRouteRequest) FrontendResponse {
	newResponse := GetCoordinatesAndInfoFromRoute(route, mode)
	newResponse.Route[len(newResponse.Route)-1].Instruction += (" <i>" + request.Destination + "</i>")
	// fmt.Printf("%# v", pretty.Formatter(newResponse))
	if len(base.Route) == 0 {
		return newResponse
	}
	// fmt.Printf("%# v", pretty.Formatter(base))
	polylineA, _ := base.OverviewPolyline.Decode()
	polylineB, _ := newResponse.OverviewPolyline.Decode()

	newPolyline := append(polylineA, polylineB...)

	return FrontendResponse{
		Route:            append(base.Route, newResponse.Route...),
		ArrivalTime:      newResponse.ArrivalTime,
		DepartureTime:    newResponse.DepartureTime,
		OverviewPolyline: maps.Polyline{Points: maps.Encode(newPolyline)},
	}
}

func ParsePlaces(clientRequest url.Values) ([]Place, Place) {

	source := Place{
		Name:         "",
		OpeningHours: "00:00-24:00",
		Time:         "0",
		PlaceID:      "",
	}

	waypoints := make([]Place, 0)

	for key, value := range clientRequest {
		fmt.Println(key, value)
		if value[0] == "undefined" || value[0] == "null" {
			continue
		}

		switch key {
		case "origin":
			source.Name = value[0]
		case "waypoints":
			var placetList []Place
			_ = json.Unmarshal([]byte(value[0]), &placetList)
			waypoints = placetList
		}
	}
	for _, place := range waypoints {
		if place.Name == source.Name {
			source = place
			break
		}
	}

	waypoints = append(waypoints)
	return waypoints, source
}

func ParseFetchPlacesRequest(c *gin.Context) GoogleCustomNearbySearchRequest {
	q := c.Request.URL.Query()
	north, _ := strconv.ParseFloat(q["north"][0], 64)
	south, _ := strconv.ParseFloat(q["south"][0], 64)
	west, _ := strconv.ParseFloat(q["west"][0], 64)
	east, _ := strconv.ParseFloat(q["east"][0], 64)
	println(north)
	latitude :=
		strconv.FormatFloat((north+south)/2.0, 'f', 6, 64) +
			"," +
			strconv.FormatFloat((west+east)/2.0, 'f', 6, 64)
	println(latitude)
	return GoogleCustomNearbySearchRequest{
		Location:   latitude,
		RankBy:     "distance",
		PlaceTypes: q["types"][0],
	}
}
