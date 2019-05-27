package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gpr95/ItineraryAdvisor/trip"
)

type place struct {
	Name         string
	OpeningHours string
	Time         string
	PlaceID      string
}

var places = []place{
	{Name: "Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours: "10:00-18:00", Time: "1h", PlaceID: "ChIJV9L_QvXMHkcR4fE_qS-WapY"},
	{Name: "Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours: "11:00-18:00", Time: "1h", PlaceID: "ChIJj1ffSl7MHkcRBfCWp3BZgGc"},
	{Name: "Muzeum Fryderyka Chopina w Warszawie", OpeningHours: "11:00-20:00", Time: "1h", PlaceID: "ChIJKbVhrFjMHkcRdbMJyIdXC34"},
	{Name: "Muzeum Uniwersytetu Warszawskiego", OpeningHours: "", Time: "1h", PlaceID: "ChIJnT24I17MHkcR8TirCHITPHM"},
	{Name: "Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours: "10:00-17:00", Time: "1h", PlaceID: "ChIJPcztF2DMHkcRNTXS6e61rqo"},
	{Name: "Muzeum Historii Akimosik", OpeningHours: "10:00-15:00", Time: "1h", PlaceID: "ChIJZ1x3xfTMHkcRAimNnQ9D9os"},
	{Name: "Muzeum Narodowe w Warszawie", OpeningHours: "10:00-18:00", Time: "1h", PlaceID: "ChIJK6UEcvfMHkcR0iGDJoZsIgc"},
	{Name: "Muzeum Wojska Polskiego", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJK6UEcvfMHkcRGDePQvvQqow"},
	{Name: "Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJ7-9EM_fMHkcRVh9fPLxfJW8"},
	{Name: "Muzeum Teatralne", OpeningHours: "", Time: "1h", PlaceID: "ChIJbQJ-qWbMHkcRrrYzTC9PLNw"},
	{Name: "um Sztuki Nowoczesnej w Warszawie Museum of Modern Art in Warsaw Oddział nad Wisłą", OpeningHours: "", Time: "1h", PlaceID: "ChIJIelSflvMHkcRRl4TMOwgVEs"},
	{Name: "Muzeum Sztuki Nowoczesnej w Warszawie - Muzeum nad Wisłą", OpeningHours: "12:00-20:00", Time: "1h", PlaceID: "ChIJi-z6GYzMHkcRhyyhcBDZfqk"},
	{Name: "Muzeum Karykatury im. E. Lipińskiego", OpeningHours: "10:00-18:00", Time: "1h", PlaceID: "ChIJ3SDb8mbMHkcRW0EyA4S7MCw"},
	{Name: "Muzeum Wódki", OpeningHours: "11:00-18:00", Time: "1h", PlaceID: "ChIJ48K19WPMHkcRaZQ0RMUYAkg"},
	{Name: "Park Miniatur Województwa Mazowieckiego", OpeningHours: "11:00-19:00", Time: "1h", PlaceID: "ChIJYdb1bGPMHkcRIJ3LWTCSNHw"},
	{Name: "Muzeum Domków dla Lalek w Warszawie", OpeningHours: "09:00-19:00", Time: "1h", PlaceID: "ChIJG7gk7ozMHkcRMsgidKA4Hnc"},
	{Name: "Muzeum Techniki i Przemysłu NOT", OpeningHours: "09:00-18:00", Time: "1h", PlaceID: "ChIJSS5pkozMHkcRJHxuJItu_ns"},
	{Name: "Muzeum Ziemi PAN", OpeningHours: "09:00-16:00", Time: "1h", PlaceID: "ChIJ8-qagfnMHkcRtHfL-VZmuoA"},
	{Name: "Muzeum Niepodległości", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJ705svGTMHkcRiL1PdybCAms"},
	{Name: "Muzeum Archidiecezji Warszawskiej", OpeningHours: "12:00-18:00", Time: "1h", PlaceID: "ChIJla37nP_MHkcRkX9wBjfhT2A"},
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

			googleResponse := trip.Route(trip.ParseFrontendRequest(context.Request.PostForm))
			response := trip.GetCoordinatesAndInfoFromRoute(googleResponse)
			// fmt.Printf("%# v", pretty.Formatter(response))

			context.JSON(200, response)
		})

		// Serve default places
		api.GET("/places", func(c *gin.Context) {
			// TEMPORARY MOCK - uncomment to use google api
			request := trip.ParseFetchPlacesRequest(c)
			places := trip.NearbySearch(request)
			c.JSON(200, places)
		})
	}

	// Start and run the server
	_ = router.Run(":8000")
}

// TODO Use trip.Dijkstra from debug.go in POST request and ask google API for every route to get path
// TODO Fetch all responses into single one and send to frontend
// TODO Add checkboxes in frontend to find all PlaceTypes in trip.ParseFetchPlacesRequest (do it in the loop - do not debug to much it costs)
// TODO Add waypointstime range (not single time) 3h 15m -> 3h - 4h

// TODO implement Our Algorithm (get proper weights for each waypoint (get lat lon and calculate distances, add places openhours and waypointtimes range
