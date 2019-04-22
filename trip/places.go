package trip

import (
	"context"
	"log"
	"strings"
	"time"

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

func Place(request GoogleCustomPlacesRequest) maps.PlaceDetailsResult {
	client := getGoogleClient()

	r := &maps.FindPlaceFromTextRequest{
		Input:     request.Input,
		InputType: lookupInputType(request.InputType),
	}

	place, err := client.FindPlaceFromText(context.Background(), r)
	check(err)


	detailPlaceRequest := &maps.PlaceDetailsRequest{
		PlaceID: place.Candidates[0].PlaceID,
	}
	placeDetail, err := client.PlaceDetails(context.Background(), detailPlaceRequest)


	//fmt.Printf("%# v", pretty.Formatter(placeDetail))
	current := time.Now()

	println("Opening hours of: " + placeDetail.Name + " at current week day : " + getOpeningHours(placeDetail, current))

	return placeDetail
}

func getFormattedAddress(place maps.PlaceDetailsResult) string{
	return place.FormattedAddress
}

func getOpeningHours(place maps.PlaceDetailsResult, departureTime time.Time) string {
	weekDay := departureTime.Weekday().String()
	println(weekDay)

	for _, day := range place.OpeningHours.WeekdayText {
		if strings.Contains(day, weekDay) {
			return strings.Replace(day, weekDay + ": ", "", 1)
		}
	}

	return ""
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
