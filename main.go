package main

func main() {
	//request := GoogleCustomRouteRequest{
	//	origin:                   "Chełm",
	//	destination:              "Warszawa",
	//	mode:                     "driving",
	//	departureTime:            "",
	//	arrivalTime:              "",
	//	waypoints:                "",
	//	language:                 "PL",
	//	region:                   "",
	//	transitMode:              "",
	//	transitRoutingPreference: "",
	//	trafficModel:             "",
	//}
	//route(request)

	requestPlaces := GoogleCustomPlacesRequest{
		input:     "Museum of Contemporary Art Australia",
		inputType: "textquery",
		fields:    "photos,formatted_address,name,rating,opening_hours,geometry",
	}

	place(requestPlaces)
}
