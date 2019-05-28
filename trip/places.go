package trip

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

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

type GoogleCustomNearbySearchRequest struct {
	Location   string
	Radius     uint
	RankBy     string
	PlaceTypes string // pipe sign separated
}

type Place struct {
	Name         string
	OpeningHours string
	Time         string
	PlaceID      string
}

type RouteStatistics struct {
	Wight float64
	Distance int
	Time time.Duration
}

func PlaceFinder(request GoogleCustomPlacesRequest) maps.PlaceDetailsResult {
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

func NearbySearch(request GoogleCustomNearbySearchRequest) []Place {
	client := getGoogleClient()

	r := &maps.NearbySearchRequest{
		Radius:   request.Radius,
		Keyword:  request.PlaceTypes,
		Language: "PL",
	}
	parseLocation(request.Location, r)
	parseRankBy(request.RankBy, r)
	parsePlaceType(request.PlaceTypes, r)

	resp, err := client.NearbySearch(context.Background(), r)
	check(err)

	places := make([]Place, 0)
	for i := 0; i < len(resp.Results); i++ {
		detailPlaceRequest := &maps.PlaceDetailsRequest{
			PlaceID:  resp.Results[i].PlaceID,
			Language: "PL",
		}
		placeDetail, err := client.PlaceDetails(context.Background(), detailPlaceRequest)
		places = append(places,
			Place{
				Name:         resp.Results[i].Name,
				OpeningHours: getOpeningHours(placeDetail, time.Now()),
				Time:         "1h",
				PlaceID:      resp.Results[i].PlaceID},
		)
		check(err)
	}

	fmt.Printf("%# v", pretty.Formatter(places))
	return places
}


func GetWightsBetweenPlaces(placesIDs []Place) map[Place]map[Place]RouteStatistics{
	client := getGoogleClient()

	distanceMatrixRequest := &maps.DistanceMatrixRequest{
		Language: "PL",
	}

	var placesIDsChnged []string

	for _, place := range placesIDs {
		if len(place.PlaceID) != 0 {
			placesIDsChnged = append(placesIDsChnged, "place_id:"+place.PlaceID)
		} else {
			placesIDsChnged = append(placesIDsChnged, place.Name)
		}

	}
	fmt.Printf("%# v", pretty.Formatter(placesIDsChnged))
	distanceMatrixRequest.Origins = placesIDsChnged
	distanceMatrixRequest.Destinations = placesIDsChnged

	distanceMatrixRequest.Mode = maps.TravelModeWalking
	distanceMatrixRequest.Units = maps.UnitsMetric

	resp, err := client.DistanceMatrix(context.Background(), distanceMatrixRequest)
	check(err)

	statistics := make(map[Place]map[Place]RouteStatistics)

	for idx, placeObject := range placesIDs {
		innerDistances := make(map[Place]RouteStatistics)

		for idxDistance, row := range resp.Rows[idx].Elements{
			if placesIDs[idxDistance] != placeObject {
				wight := float64(row.Distance.Meters) +
					parseOpenHourToWight(placesIDs[idxDistance].OpeningHours, row.Distance.Meters) +
					parseTimeToWight(placesIDs[idxDistance].Time, row.Distance.Meters)
				time := row.Duration
				statistics := RouteStatistics{Distance: row.Distance.Meters, Time: time, Wight: wight}

				innerDistances[placesIDs[idxDistance]] = statistics
			}
		}
		statistics[placeObject] = innerDistances
	}

	fmt.Printf("%# v", pretty.Formatter(statistics))

	return statistics
}

func parseTimeToWight(stayingTime string, distance int) float64 {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(stayingTime, "")

	hours, _ := strconv.Atoi(processedString)
	fmt.Printf("%s -> %f, distance: %d\n ", stayingTime, -(float64(hours)/24.0)*float64(distance), distance)
	return -(float64(hours) / 24.0) * float64(distance)
}

func parseOpenHourToWight(openingHours string, distance int) float64 {
	closingTime := strings.Split(openingHours, "-")[1]
	closingHourStr := strings.Split(closingTime, ":")[0]
	closingHour, _ := strconv.Atoi(closingHourStr)
	fmt.Printf("%s -> %f, distance: %d\n ", openingHours, -(float64(closingHour)/24.0)*float64(distance), distance)
	return -(float64(closingHour) / 24.0) * float64(distance)
}

func getFormattedAddress(place maps.PlaceDetailsResult) string {
	return place.FormattedAddress
}

func getOpeningHours(place maps.PlaceDetailsResult, departureTime time.Time) string {
	weekDay := departureTime.Weekday().String()
	weekdayMap := map[string]string{
		"Monday":    "poniedziałek",
		"Tuesday":   "wtorek",
		"Wednesday": "środa",
		"Thursday":  "czwartek",
		"Friday":    "piątek",
		"Saturday":  "sobota",
		"Sunday":    "niedziela",
	}
	if place.OpeningHours == nil || len(place.OpeningHours.WeekdayText) == 0 {
		return ""
	}

	//fmt.Printf("%# v", pretty.Formatter(place.OpeningHours))
	for _, day := range place.OpeningHours.WeekdayText {
		if strings.Contains(day, weekdayMap[weekDay]) {
			openingHours := strings.Replace(day, weekdayMap[weekDay]+": ", "", 1)
			replacedDash := strings.Replace(openingHours, "–", "-", 1)
			if openingHours == "Zamknięte" {
				return "00:00-24:00"
			}
			return replacedDash
		}
	}

	return "00:00-24:00"
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
