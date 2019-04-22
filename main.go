package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gpr95/ItineraryAdvisor/trip"
	"github.com/kr/pretty"
)

type place struct {
	Name         string
	OpeningHours string
	Time         string
}

var places = []place{
	{Name: "Grób nieznanego żołnierza Warszawa", OpeningHours: "10:00-20:00", Time: "1h"},
	{Name: "Muzeum więzienia Pawiak Warszawa", OpeningHours: "11:00-20:00", Time: "7h"},
	{Name: "Kino Luna Warszawa", OpeningHours: "13:00-20:00", Time: "6h"},
	{Name: "Fort Legionów Warszawa", OpeningHours: "12:00-20:00", Time: "3h"},
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
			fmt.Printf("%# v", pretty.Formatter(places))
			c.JSON(200, places)
		})
	}

	// Start and run the server
	_ = router.Run(":8000")
}