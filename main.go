package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// request := GoogleCustomRouteRequest{
	// 	origin:                   "",
	// 	destination:              "",
	// 	mode:                     "",
	// 	departureTime:            "",
	// 	arrivalTime:              "",
	// 	waypoints:                "",
	// 	language:                 "PL",
	// 	region:                   "",
	// 	transitMode:              "",
	// 	transitRoutingPreference: "",
	// 	trafficModel:             "",
	// }
	//route(request)

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
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.POST("/form-submit-url", func(c *gin.Context) {
			c.Request.ParseMultipartForm(1000)
			for key, value := range c.Request.PostForm {
				fmt.Println(key, value)
			}
		})
	}

	// Start and run the server
	router.Run(":8000")

}
