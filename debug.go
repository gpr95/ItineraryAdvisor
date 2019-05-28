// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main contains a simple command line tool for Places API Text Search
// Documentation: https://developers.google.com/places/web-service/search#TextSearchRequests
package main

import (
	"fmt"

	"github.com/gpr95/ItineraryAdvisor/trip"
	"github.com/kr/pretty"
)

var places2 = []trip.Place{
	{Name: "Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours: "10:00-18:00", Time: "1h", PlaceID: "ChIJV9L_QvXMHkcR4fE_qS-WapY"},
	{Name: "Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours: "11:00-18:00", Time: "1h", PlaceID: "ChIJj1ffSl7MHkcRBfCWp3BZgGc"},
	{Name: "Muzeum Fryderyka Chopina w Warszawie", OpeningHours: "11:00-20:00", Time: "1h", PlaceID: "ChIJKbVhrFjMHkcRdbMJyIdXC34"},
	{Name: "Muzeum Uniwersytetu Warszawskiego", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJnT24I17MHkcR8TirCHITPHM"},
	{Name: "Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours: "10:00-17:00", Time: "1h", PlaceID: "ChIJPcztF2DMHkcRNTXS6e61rqo"},
	{Name: "Muzeum Historii Akimosik", OpeningHours: "10:00-15:00", Time: "1h", PlaceID: "ChIJZ1x3xfTMHkcRAimNnQ9D9os"},
	{Name: "Muzeum Narodowe w Warszawie", OpeningHours: "10:00-18:00", Time: "1h", PlaceID: "ChIJK6UEcvfMHkcR0iGDJoZsIgc"},
	{Name: "Muzeum Wojska Polskiego", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJK6UEcvfMHkcRGDePQvvQqow"},
	{Name: "Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJ7-9EM_fMHkcRVh9fPLxfJW8"},
	{Name: "Muzeum Teatralne", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJbQJ-qWbMHkcRrrYzTC9PLNw"},
}

func main() {
	path := trip.FindItinerary(places2, trip.Place{Name: "Muzeum Teatralne", OpeningHours: "00:00-00:00", Time: "1h", PlaceID: "ChIJbQJ-qWbMHkcRrrYzTC9PLNw"})
	googleRequestsList := trip.ParseItineraryToGoogleRequests(path)

	frontendResponse := trip.FrontendResponse{}
	for _, googleRequest := range googleRequestsList {
		route := trip.Route(googleRequest)
		frontendResponse = trip.AppendGoogleResponse(frontendResponse, route)
	}
	fmt.Printf("%# v", pretty.Formatter(frontendResponse))
}
