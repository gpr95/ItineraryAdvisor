package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gpr95/ItineraryAdvisor/trip"
	"github.com/kr/pretty"
)

type place struct {
	Name         string
	OpeningHours string
	Time         string
	PlaceID      string
}

var places = []trip.Place{
	{Name: "Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours: "10:00-18:00", Time: "1h", PlaceID: "ChIJV9L_QvXMHkcR4fE_qS-WapY"},
	{Name: "Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours: "11:00-18:00", Time: "1h", PlaceID: "ChIJj1ffSl7MHkcRBfCWp3BZgGc"},
	{Name: "Muzeum Fryderyka Chopina w Warszawie", OpeningHours: "11:00-20:00", Time: "1h", PlaceID: "ChIJKbVhrFjMHkcRdbMJyIdXC34"},
	{Name: "Muzeum Uniwersytetu Warszawskiego", OpeningHours: "00:00-24:00", Time: "1h", PlaceID: "ChIJnT24I17MHkcR8TirCHITPHM"},
	{Name: "Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours: "10:00-17:00", Time: "1h", PlaceID: "ChIJPcztF2DMHkcRNTXS6e61rqo"},
	{Name: "Muzeum Historii Akimosik", OpeningHours: "10:00-15:00", Time: "1h", PlaceID: "ChIJZ1x3xfTMHkcRAimNnQ9D9os"},
	{Name: "Muzeum Narodowe w Warszawie", OpeningHours: "10:00-18:00", Time: "1h", PlaceID: "ChIJK6UEcvfMHkcR0iGDJoZsIgc"},
	{Name: "Muzeum Wojska Polskiego", OpeningHours: "00:00-24:00", Time: "1h", PlaceID: "ChIJK6UEcvfMHkcRGDePQvvQqow"},
	{Name: "Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours: "00:00-24:00", Time: "1h", PlaceID: "ChIJ7-9EM_fMHkcRVh9fPLxfJW8"},
	{Name: "Muzeum Teatralne", OpeningHours: "00:00-24:00", Time: "1h", PlaceID: "ChIJbQJ-qWbMHkcRrrYzTC9PLNw"},
}

func main() {
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		// Serve form for finding route
		api.POST("/form-submit-url", func(context *gin.Context) {
			// parses a request body as multipart/form-data.
			err := context.Request.ParseMultipartForm(1000)
			if err != http.ErrNotMultipart {
				println("error on parse multipart form map: %v", err)
			}

			// Map request objects to places
			waypoints, source := trip.ParsePlaces(context.Request.PostForm)
			fmt.Printf("%# v", pretty.Formatter(waypoints))
			fmt.Printf("%# v", pretty.Formatter(source))

			// Get google request from request
			googleRequest := trip.ParseFrontendRequest(context.Request.PostForm)

			// Run algorithm
			path, distanceSum, durationSum, arrivalTime := trip.FindItinerary(waypoints, source,
				trip.ParseDateToCustomTimeString(googleRequest.DepartureTime), googleRequest.Mode)
			distanceParsed := strconv.Itoa(distanceSum) + " m"
			fmt.Printf("%# v", pretty.Formatter(path))

			// Map solution to google request (to draw on map)
			googleRequestsList := trip.ParseItineraryToGoogleRequests(path)
			fmt.Printf("%# v", pretty.Formatter(googleRequestsList))

			// Append all rotues
			frontendResponse := trip.FrontendResponse{}
			for _, googleRequest := range googleRequestsList {
				frontendResponse = trip.AppendGoogleResponse(frontendResponse, trip.Route(googleRequest), googleRequest.Mode, googleRequest)
			}
			if trip.CustomTimeToMinutes(arrivalTime) > trip.CustomTimeToMinutes(trip.ParseDateToCustomTimeString(googleRequest.ArrivalTime)) {
				frontendResponse.Distance = "Not possible"
				frontendResponse.Duration = time.Duration(0)
			} else {
				frontendResponse.Distance = distanceParsed
				frontendResponse.Duration = durationSum
			}
			fmt.Printf("%# v", pretty.Formatter(frontendResponse))

			// Send response
			context.JSON(200, frontendResponse)
		})

		// Serve places for given coordinates
		api.GET("/places", func(c *gin.Context) {
			// TODO TEMPORARY MOCK - uncomment to use google api
			//request := trip.ParseFetchPlacesRequest(c)
			//places := trip.NearbySearch(request)
			c.JSON(200, places)
		})
	}

	// Start and run the server
	_ = router.Run(":8000")
}

// TODO check solution If place is Open when we get there
// TODO
