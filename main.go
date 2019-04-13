package main

import (
	"fmt"
	"net/url"

	"github.com/gpr95/ItineraryAdvisor/trip"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Waypoint struct {
	Lat string
	Lng string
}

func parseRequest(clientRequest url.Values, googleRequest *trip.GoogleCustomRouteRequest) {
	for key, value := range clientRequest {
		fmt.Println(key, value)
		if value[0] == "undefined" {
			continue
		}

		switch key {
		case "origin":
			googleRequest.Origin = value[0]
		case "destination":
			googleRequest.Destination = value[0]
		case "mode":
			googleRequest.Mode = value[0]
		case "departure":
			googleRequest.DepartureTime = value[0]
		case "arrival":
			googleRequest.ArrivalTime = value[0]
		}
	}

}

func parseResponse(googleResponse []string) {

}

func main() {

	request := trip.GoogleCustomRouteRequest{
		Origin:                   "",
		Destination:              "",
		Mode:                     "",
		DepartureTime:            "",
		ArrivalTime:              "",
		Waypoints:                "",
		Language:                 "PL",
		Region:                   "",
		TransitMode:              "",
		TransitRoutingPreference: "",
		TrafficModel:             "",
	}

	// requestPlaces := GoogleCustomPlacesRequest{
	// 	input:     "Museum of Contemporary Art Australia",
	// 	inputType: "textquery",
	// 	fields:    "photos,formatted_address,name,rating,opening_hours,geometry",
	// }

	// place(requestPlaces)

	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.POST("/form-submit-url", func(c *gin.Context) {
			c.Request.ParseMultipartForm(1000)
			parseRequest(c.Request.PostForm, &request)
			fmt.Println(request)
			trip.Route(request)
		})
	}

	// Start and run the server
	router.Run(":8000")

}
