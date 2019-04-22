package trip

import (
	"context"
	"fmt"
	"log"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

//Input       "The text Input specifying which place to search for (for example, a name, address, or phone number)."
//InputType   "The type of Input. This can be one of either textquery or phonenumber."
//Fields      "Comma seperated list of Fields"
type GoogleCustomPlacesRequest struct {
	Input     string
	InputType string
	Fields    string
}

func Place(request GoogleCustomPlacesRequest) {
	client := getGoogleClient()

	r := &maps.FindPlaceFromTextRequest{
		Input:     request.Input,
		InputType: lookupInputType(request.InputType),
	}

	place, err := client.FindPlaceFromText(context.Background(), r)
	check(err)

	fmt.Printf("%# v", pretty.Formatter(place))
	fmt.Printf("%# v", pretty.Formatter(place.Candidates[0].Geometry))
}

func lookupInputType(inputType string) maps.FindPlaceFromTextInputType {
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
