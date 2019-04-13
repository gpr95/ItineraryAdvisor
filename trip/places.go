package trip

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

//Input                     "Input", "", "The text Input specifying which place to search for (for example, a name, address, or phone number)."
//InputType                 "InputType", "", "The type of Input. This can be one of either textquery or phonenumber."
//Fields                    "Fields", "", "Comma seperated list of Fields"
type GoogleCustomPlacesRequest struct {
	Input     string
	InputType string
	Fields    string
}

func Place(request GoogleCustomPlacesRequest) {
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
		Input:     request.Input,
		InputType: parseInputType(request.InputType),
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
		log.Fatalf("Unknown Input type '%s'", inputType)
	}
	return it
}
