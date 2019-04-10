package main

import (
	"context"
	"fmt"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
	"log"
	"os"
)

//input                     "input", "", "The text input specifying which place to search for (for example, a name, address, or phone number)."
//inputType                 "inputType", "", "The type of input. This can be one of either textquery or phonenumber."
//fields                    "fields", "", "Comma seperated list of Fields"
type GoogleCustomPlacesRequest struct {
	input     string
	inputType string
	fields    string
}

func place(request GoogleCustomPlacesRequest) {
	configuration, _ := getSecrets()

	var client *maps.Client
	var err error
	if configuration.GoogleAPIKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(configuration.GoogleAPIKey), maps.WithRateLimit(2))
	} else {
		_, err = fmt.Fprintln(os.Stderr, "Please specify an API Key.")
		os.Exit(2)
	}
	check(err)

	r := &maps.FindPlaceFromTextRequest{
		Input:     request.input,
		InputType: parseInputType(request.inputType),
	}

	place, err := client.FindPlaceFromText(context.Background(), r)
	check(err)

	fmt.Printf("%# v", pretty.Formatter(place))
	fmt.Printf("%# v", pretty.Formatter(place.Candidates[0].Geometry))
}

func parseInputType(inputType string) maps.FindPlaceFromTextInputType {
	var it maps.FindPlaceFromTextInputType
	switch inputType {
	case "textquery":
		it = maps.FindPlaceFromTextInputTypeTextQuery
	case "phonenumber":
		it = maps.FindPlaceFromTextInputTypePhoneNumber
	default:
		log.Fatalf("Unknown input type '%s'", inputType)
	}
	return it
}
