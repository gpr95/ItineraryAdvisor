package trip

import (
	"context"
	"fmt"
	"github.com/kr/pretty"
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

type GoogleCustomNearbySearchRequest struct {
	location string
	radius uint
	keyword string
	rankBy string
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

func NearbySearch(request GoogleCustomNearbySearchRequest) string {
	client := getGoogleClient()

	r := &maps.NearbySearchRequest{
		Radius:    request.radius,
		Keyword:   request.keyword,
		Language:  "PL",
	}
	parseLocation(request.location, r)
	parseRankBy(request.rankBy, r)

	resp, err := client.NearbySearch(context.Background(), r)
	check(err)
	fmt.Printf("%# v", pretty.Formatter(resp))
	return ""
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

func parseLocation(location string, r *maps.NearbySearchRequest) {
	if location != "" {
		l, err := maps.ParseLatLng(location)
		check(err)
		r.Location = &l
	}
}


func parsePriceLevel(priceLevel string) maps.PriceLevel {
	switch priceLevel {
	case "0":
		return maps.PriceLevelFree
	case "1":
		return maps.PriceLevelInexpensive
	case "2":
		return maps.PriceLevelModerate
	case "3":
		return maps.PriceLevelExpensive
	case "4":
		return maps.PriceLevelVeryExpensive
	default:
		log.Fatalf(fmt.Sprintf("Unknown price level: '%s'", priceLevel))
	}
	return maps.PriceLevelFree
}

func parsePriceLevels(minPrice string, maxPrice string, r *maps.NearbySearchRequest) {
	if minPrice != "" {
		r.MinPrice = parsePriceLevel(minPrice)
	}

	if maxPrice != "" {
		r.MaxPrice = parsePriceLevel(minPrice)
	}
}

func parseRankBy(rankBy string, r *maps.NearbySearchRequest) {
	switch rankBy {
	case "prominence":
		r.RankBy = maps.RankByProminence
		return
	case "distance":
		r.RankBy = maps.RankByDistance
		return
	case "":
		return
	default:
		log.Fatalf(fmt.Sprintf("Unknown rank by: \"%v\"", rankBy))
	}
}

func parsePlaceType(placeType string, r *maps.NearbySearchRequest) {
	if placeType != "" {
		t, err := maps.ParsePlaceType(placeType)
		if err != nil {
			log.Fatalf(fmt.Sprintf("Unknown place type \"%v\"", placeType))
		}

		r.Type = t
	}
}