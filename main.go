package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gpr95/ItineraryAdvisor/trip"
	"github.com/kr/pretty"
	// "github.com/kr/pretty"
)

type place struct {
	Name         string
	OpeningHours string
	Time         string
}

var places = []place{
	place{Name: "Grób nieznanego żołnierza Warszawa", OpeningHours: "10:00-20:00", Time: "1h"},
	place{Name: "Muzeum więzienia Pawiak Warszawa", OpeningHours: "11:00-20:00", Time: "7h"},
	place{Name: "Kino Luna Warszawa", OpeningHours: "13:00-20:00", Time: "6h"},
	place{Name: "Fort Legionów Warszawa", OpeningHours: "12:00-20:00", Time: "3h"},
}

func main() {

	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.POST("/form-submit-url", func(c *gin.Context) {
			c.Request.ParseMultipartForm(1000)
			request := trip.ParseFrontendRequest(c.Request.PostForm)
			googleResponse := trip.Route(request)
			response := trip.GetCoordinatesAndInfoFromRoute(googleResponse)
			// fmt.Printf("%# v", pretty.Formatter(response))

			c.JSON(200, response)
		})
		api.GET("/places", func(c *gin.Context) {
			fmt.Printf("%# v", pretty.Formatter(places))
			c.JSON(200, places)
		})
	}

	// Start and run the server
	router.Run(":8000")
}
