package main

func main() {
	request := GoogleCustomRequest{
		origin:                   "Chełm",
		destination:              "Warszawa",
		mode:                     "driving",
		departureTime:            "",
		arrivalTime:              "",
		waypoints:                "",
		language:                 "PL",
		region:                   "",
		transitMode:              "",
		transitRoutingPreference: "",
		trafficModel:             "",
	}
	route(request)
}
