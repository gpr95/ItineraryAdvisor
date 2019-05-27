package trip

import (
	"bytes"
	"container/heap"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"sync"
)

func CreateJSONGraph(fileName string, graphName string, waypoints []Place)  {
	//TODO temporary MOCK
	//wayPointsWithWights := GetWightsBetweenPlaces(waypoints)
	wayPointsWithWights := map[Place]map[Place]float64{
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:30.625, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:91.5, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:0, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:271.0416666666667, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:1390.5416666666667, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:1199.8333333333333, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:779.125, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:130, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:123, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:379.3333333333333},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:1090.5833333333333, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:0, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:1104, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:107.08333333333334, {Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:206.66666666666669, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:114.62499999999993, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:168, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:208.75, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:893.1666666666666, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:1386.7083333333333},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:592.25, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:268.75, {Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:271.875, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:94.125, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:384, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:0, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:1390.5416666666667, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:424.99999999999994, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:38.125, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:2169.6666666666665},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:195, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:190.125, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:196.74999999999997, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:482.3333333333333, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:2169.6666666666665, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:236.875, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:440.4166666666667, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:1978.9583333333333, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:0, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:756.125},
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:99.375, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:171.33333333333334, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:237.5, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:1045.5416666666667, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:1089.625, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:0, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:598, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:102.49999999999999, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:1236.25, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:77.87499999999993},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:140.875, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:100.74999999999997, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:1059.9166666666667, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:99.375, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:73.125, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:240.41666666666669, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:1250.625, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:920, {Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:0, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:330.6666666666667},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:121.875, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:0, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:701.5, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:246.99999999999997, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:1006.25, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:129.79166666666666, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:305.6666666666667, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:179.58333333333334, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:715.875, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:1480.625},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:85.41666666666667, {Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:83.95833333333334, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:1438.4583333333333, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:1629.1666666666667, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:754.2083333333334, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:123.49999999999997, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:471.5, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:0, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:224, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:322.9166666666667},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:240.41666666666669, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:387.49999999999994, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:175.375, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:2025.9166666666667, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:0, {Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:448.5, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:237.5, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:108.49999999999997, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:1246.7916666666667, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:334},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:0, {Name:"Muzeum Teatralne", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:1978.9583333333333, {Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:227.29166666666669, {Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}:230.41666666666669, {Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:375.24999999999983, {Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:310.6666666666667, {Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:97.5, {Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:131.25, {Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:1199.8333333333333, {Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-00:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:592.25},
	}

	mapWithGraphName := make(map[string]map[string]map[string]float64)
	wholeMap := make(map[string]map[string]float64)
	for i := 0; i < len(waypoints) ; i++  {

		otherNodes := make(map[string]float64)

		for k, v := range wayPointsWithWights[waypoints[i]] {
			otherNodes[k.Name] = float64(v)
		}

		wholeMap[waypoints[i].Name] = otherNodes
	}
	mapWithGraphName[graphName] = wholeMap
	wholeMapJSON, _ := json.Marshal(mapWithGraphName)
	fmt.Print(string(wholeMapJSON))

	writeToFile(fileName, string(wholeMapJSON))
}

func writeToFile(fileName string, data string) {
	dataBytes := []byte(data)
	err := ioutil.WriteFile(fileName, dataBytes, 0644)
	check(err)
}

type nodeDistance struct {
	id       ID
	distance float64
}

// container.Heap's Interface needs sort.Interface, Push, Pop to be implemented

// nodeDistanceHeap is a min-heap of nodeDistances.
type nodeDistanceHeap []nodeDistance

func (h nodeDistanceHeap) Len() int           { return len(h) }
func (h nodeDistanceHeap) Less(i, j int) bool { return h[i].distance < h[j].distance } // Min-Heap
func (h nodeDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(nodeDistance))
}

func (h *nodeDistanceHeap) Pop() interface{} {
	heapSize := len(*h)
	lastNode := (*h)[heapSize-1]
	*h = (*h)[0 : heapSize-1]
	return lastNode
}

func (h *nodeDistanceHeap) updateDistance(id ID, val float64) {
	for i := 0; i < len(*h); i++ {
		if (*h)[i].id == id {
			(*h)[i].distance = val
			break
		}
	}
}


func Dijkstra(g Graph, source, target ID) ([]ID, map[ID]float64, error) {
	// let Q be a priority queue
	minHeap := &nodeDistanceHeap{}

	// distance[source] = 0
	distance := make(map[ID]float64)
	distance[source] = 0.0

	// for each vertex v in G:
	for id := range g.GetNodes() {
		// if v ≠ source:
		if id != source {
			// distance[v] = ∞
			distance[id] = math.MaxFloat64

			// prev[v] = undefined
			// prev[v] = ""
		}

		// Q.add_with_priority(v, distance[v])
		nds := nodeDistance{}
		nds.id = id
		nds.distance = distance[id]

		heap.Push(minHeap, nds)
	}

	heap.Init(minHeap)
	prev := make(map[ID]ID)

	// while Q is not empty:
	for minHeap.Len() != 0 {

		// u = Q.extract_min()
		u := heap.Pop(minHeap).(nodeDistance)

		// if u == target:
		if u.id == target {
			break
		}

		// for each child vertex v of u:
		cmap, err := g.GetTargets(u.id)
		if err != nil {
			return nil, nil, err
		}
		for v := range cmap {

			// alt = distance[u] + weight(u, v)
			weight, err := g.GetWeight(u.id, v)
			if err != nil {
				return nil, nil, err
			}
			alt := distance[u.id] + weight

			// if distance[v] > alt:
			if distance[v] > alt {

				// distance[v] = alt
				distance[v] = alt

				// prev[v] = u
				prev[v] = u.id

				// Q.decrease_priority(v, alt)
				minHeap.updateDistance(v, alt)
			}
		}
		heap.Init(minHeap)
	}

	// path = []
	path := []ID{}

	// u = target
	u := target

	// while prev[u] is defined:
	for {
		if _, ok := prev[u]; !ok {
			break
		}
		// path.push_front(u)
		temp := make([]ID, len(path)+1)
		temp[0] = u
		copy(temp[1:], path)
		path = temp

		// u = prev[u]
		u = prev[u]
	}

	// add the source
	temp := make([]ID, len(path)+1)
	temp[0] = source
	copy(temp[1:], path)
	path = temp

	return path, distance, nil
}

// ID is unique identifier.
type ID interface {
	// String returns the string ID.
	String() string
}

type StringID string

func (s StringID) String() string {
	return string(s)
}

// Node is vertex. The ID must be unique within the graph.
type Node interface {
	// ID returns the ID.
	ID() ID
	String() string
}

type node struct {
	id string
}

var nodeCnt uint64

func NewNode(id string) Node {
	return &node{
		id: id,
	}
}

func (n *node) ID() ID {
	return StringID(n.id)
}

func (n *node) String() string {
	return n.id
}

// Edge connects between two Nodes.
type Edge interface {
	Source() Node
	Target() Node
	Weight() float64
	String() string
}

// edge is an Edge from Source to Target.
type edge struct {
	src Node
	tgt Node
	wgt float64
}

func NewEdge(src, tgt Node, wgt float64) Edge {
	return &edge{
		src: src,
		tgt: tgt,
		wgt: wgt,
	}
}

func (e *edge) Source() Node {
	return e.src
}

func (e *edge) Target() Node {
	return e.tgt
}

func (e *edge) Weight() float64 {
	return e.wgt
}

func (e *edge) String() string {
	return fmt.Sprintf("%s -- %.3f -→ %s\n", e.src, e.wgt, e.tgt)
}

type EdgeSlice []Edge

func (e EdgeSlice) Len() int           { return len(e) }
func (e EdgeSlice) Less(i, j int) bool { return e[i].Weight() < e[j].Weight() }
func (e EdgeSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }


// Graph describes the methods of graph operations.
// It assumes that the identifier of a Node is unique.
// And weight values is float64.
type Graph interface {
	// Init initializes a Graph.
	Init()

	// GetNodeCount returns the total number of nodes.
	GetNodeCount() int

	// GetNode finds the Node. It returns nil if the Node
	// does not exist in the graph.
	GetNode(id ID) Node

	// GetNodes returns a map from node ID to
	// empty struct value. Graph does not allow duplicate
	// node ID or name.
	GetNodes() map[ID]Node

	// AddNode adds a node to a graph, and returns false
	// if the node already existed in the graph.
	AddNode(nd Node) bool

	// DeleteNode deletes a node from a graph.
	// It returns true if it got deleted.
	// And false if it didn't get deleted.
	DeleteNode(id ID) bool

	// AddEdge adds an edge from nd1 to nd2 with the weight.
	// It returns error if a node does not exist.
	AddEdge(id1, id2 ID, weight float64) error

	// ReplaceEdge replaces an edge from id1 to id2 with the weight.
	ReplaceEdge(id1, id2 ID, weight float64) error

	// DeleteEdge deletes an edge from id1 to id2.
	DeleteEdge(id1, id2 ID) error

	// GetWeight returns the weight from id1 to id2.
	GetWeight(id1, id2 ID) (float64, error)

	// GetSources returns the map of parent Nodes.
	// (Nodes that come towards the argument vertex.)
	GetSources(id ID) (map[ID]Node, error)

	// GetTargets returns the map of child Nodes.
	// (Nodes that go out of the argument vertex.)
	GetTargets(id ID) (map[ID]Node, error)

	// String describes the Graph.
	String() string
}

// graph is an internal default graph type that
// implements all methods in Graph interface.
type graph struct {
	mu sync.RWMutex // guards the following

	// idToNodes stores all nodes.
	idToNodes map[ID]Node

	// nodeToSources maps a Node identifer to sources(parents) with edge weights.
	nodeToSources map[ID]map[ID]float64

	// nodeToTargets maps a Node identifer to targets(children) with edge weights.
	nodeToTargets map[ID]map[ID]float64
}

// newGraph returns a new graph.
func newGraph() *graph {
	return &graph{
		idToNodes:     make(map[ID]Node),
		nodeToSources: make(map[ID]map[ID]float64),
		nodeToTargets: make(map[ID]map[ID]float64),
		//
		// without this
		// panic: assignment to entry in nil map
	}
}

// NewGraph returns a new graph.
func NewGraph() Graph {
	return newGraph()
}

func (g *graph) Init() {
	// (X) g = newGraph()
	// this only updates the pointer
	//
	//
	// (X) *g = *newGraph()
	// assignment copies lock value

	g.idToNodes = make(map[ID]Node)
	g.nodeToSources = make(map[ID]map[ID]float64)
	g.nodeToTargets = make(map[ID]map[ID]float64)
}

func (g *graph) GetNodeCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return len(g.idToNodes)
}

func (g *graph) GetNode(id ID) Node {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.idToNodes[id]
}

func (g *graph) GetNodes() map[ID]Node {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.idToNodes
}

func (g *graph) unsafeExistID(id ID) bool {
	_, ok := g.idToNodes[id]
	return ok
}

func (g *graph) AddNode(nd Node) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.unsafeExistID(nd.ID()) {
		return false
	}

	id := nd.ID()
	g.idToNodes[id] = nd
	return true
}

func (g *graph) DeleteNode(id ID) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id) {
		return false
	}

	delete(g.idToNodes, id)

	delete(g.nodeToTargets, id)
	for _, smap := range g.nodeToTargets {
		delete(smap, id)
	}

	delete(g.nodeToSources, id)
	for _, smap := range g.nodeToSources {
		delete(smap, id)
	}

	return true
}

func (g *graph) AddEdge(id1, id2 ID, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id1) {
		return fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.nodeToTargets[id1]; ok {
		if v, ok2 := g.nodeToTargets[id1][id2]; ok2 {
			g.nodeToTargets[id1][id2] = v + weight
		} else {
			g.nodeToTargets[id1][id2] = weight
		}
	} else {
		tmap := make(map[ID]float64)
		tmap[id2] = weight
		g.nodeToTargets[id1] = tmap
	}
	if _, ok := g.nodeToSources[id2]; ok {
		if v, ok2 := g.nodeToSources[id2][id1]; ok2 {
			g.nodeToSources[id2][id1] = v + weight
		} else {
			g.nodeToSources[id2][id1] = weight
		}
	} else {
		tmap := make(map[ID]float64)
		tmap[id1] = weight
		g.nodeToSources[id2] = tmap
	}

	return nil
}

func (g *graph) ReplaceEdge(id1, id2 ID, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id1) {
		return fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.nodeToTargets[id1]; ok {
		g.nodeToTargets[id1][id2] = weight
	} else {
		tmap := make(map[ID]float64)
		tmap[id2] = weight
		g.nodeToTargets[id1] = tmap
	}
	if _, ok := g.nodeToSources[id2]; ok {
		g.nodeToSources[id2][id1] = weight
	} else {
		tmap := make(map[ID]float64)
		tmap[id1] = weight
		g.nodeToSources[id2] = tmap
	}
	return nil
}

func (g *graph) DeleteEdge(id1, id2 ID) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id1) {
		return fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.nodeToTargets[id1]; ok {
		if _, ok := g.nodeToTargets[id1][id2]; ok {
			delete(g.nodeToTargets[id1], id2)
		}
	}
	if _, ok := g.nodeToSources[id2]; ok {
		if _, ok := g.nodeToSources[id2][id1]; ok {
			delete(g.nodeToSources[id2], id1)
		}
	}
	return nil
}

func (g *graph) GetWeight(id1, id2 ID) (float64, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.unsafeExistID(id1) {
		return 0, fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return 0, fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.nodeToTargets[id1]; ok {
		if v, ok := g.nodeToTargets[id1][id2]; ok {
			return v, nil
		}
	}
	return 0.0, fmt.Errorf("there is no edge from %s to %s", id1, id2)
}

func (g *graph) GetSources(id ID) (map[ID]Node, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.unsafeExistID(id) {
		return nil, fmt.Errorf("%s does not exist in the graph.", id)
	}

	rs := make(map[ID]Node)
	if _, ok := g.nodeToSources[id]; ok {
		for n := range g.nodeToSources[id] {
			rs[n] = g.idToNodes[n]
		}
	}
	return rs, nil
}

func (g *graph) GetTargets(id ID) (map[ID]Node, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.unsafeExistID(id) {
		return nil, fmt.Errorf("%s does not exist in the graph.", id)
	}

	rs := make(map[ID]Node)
	if _, ok := g.nodeToTargets[id]; ok {
		for n := range g.nodeToTargets[id] {
			rs[n] = g.idToNodes[n]
		}
	}
	return rs, nil
}

func (g *graph) String() string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	buf := new(bytes.Buffer)
	for id1, nd1 := range g.idToNodes {
		nmap, _ := g.GetTargets(id1)
		for id2, nd2 := range nmap {
			weight, _ := g.GetWeight(id1, id2)
			fmt.Fprintf(buf, "%s -- %.3f -→ %s\n", nd1, weight, nd2)
		}
	}
	return buf.String()
}

func NewGraphFromJSON(rd io.Reader, graphID string) (Graph, error) {
	js := make(map[string]map[string]map[string]float64)
	dec := json.NewDecoder(rd)
	for {
		if err := dec.Decode(&js); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	if _, ok := js[graphID]; !ok {
		return nil, fmt.Errorf("%s does not exist", graphID)
	}
	gmap := js[graphID]

	g := newGraph()
	for id1, mm := range gmap {
		nd1 := g.GetNode(StringID(id1))
		if nd1 == nil {
			nd1 = NewNode(id1)
			if ok := g.AddNode(nd1); !ok {
				return nil, fmt.Errorf("%s already exists", nd1)
			}
		}
		for id2, weight := range mm {
			nd2 := g.GetNode(StringID(id2))
			if nd2 == nil {
				nd2 = NewNode(id2)
				if ok := g.AddNode(nd2); !ok {
					return nil, fmt.Errorf("%s already exists", nd2)
				}
			}
			g.ReplaceEdge(nd1.ID(), nd2.ID(), weight)
		}
	}

	return g, nil
}