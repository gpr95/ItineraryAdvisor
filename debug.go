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
	"os"
	"strings"
)

func main() {
	fileName := "testfile"
	graphName := "testGraph"
	trip.CreateJSONGraph(fileName, graphName, []string{"A", "B", "C", "D"}, []string{"3h 15m", "3h 15m", "3h 15m", "3h 15m"})

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := trip.NewGraphFromJSON(f, graphName)
	if err != nil {
		panic(err)
	}
	path, distance, err := trip.Dijkstra(g, trip.StringID("A"), trip.StringID("D"))
	if err != nil {
		panic(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}

	fmt.Println(strings.Join(ts, " → "))
	fmt.Println(distance[trip.StringID("T")])
	fmt.Println("testGraph:", strings.Join(ts, " → "))
}
