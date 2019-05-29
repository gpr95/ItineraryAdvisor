package trip

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type TransportStatistics struct {
	TransportType string
	TimeDuration time.Duration
	DistanceMeters int
}

func FindItinerary(waypoints []Place, source Place, departure string, modes []string) (map[Place]TransportStatistics, int, time.Duration, string){
	wayPointsWithWights := GetWightsBetweenPlaces(waypoints, modes)
	//wayPointsWithWights := MOCK_MAP

	path := make(map[Place]TransportStatistics)

	distanceSum := 0
	durationSum := time.Duration(0)
	initNode := source
	// Remove initNode from destinations of other nodes
	for _, destinations := range wayPointsWithWights {
		delete(destinations, initNode)
	}
	path[source] = TransportStatistics{chooseTransportBetweenPlaces(initNode, initNode), time.Duration(0), 0}
	for i := 0; i < len(waypoints); i++  {
		nextNode := initNode
		wight := math.MaxFloat64
		statistics := TransportStatistics{chooseTransportBetweenPlaces(initNode, nextNode), time.Duration(0), 0}
		for k, v := range wayPointsWithWights[initNode] {
			if v.Wight < wight {
				wight = v.Wight
				userSittingInPlaceTime, _ := time.ParseDuration(k.Time)
				statistics = TransportStatistics{v.Transport, v.Time + userSittingInPlaceTime, v.Distance}
				nextNode = k
			}
		}

		openingHours, closeHours := OpeningHoursToMinutes(nextNode.OpeningHours)
		if nextNode == initNode && CustomTimeToMinutes(departure) > closeHours && CustomTimeToMinutes(departure) < openingHours {
			log.Fatalf("fatal error: Algorithm not working...")
		}
		if nextNode != initNode {
			distanceSum = distanceSum + statistics.DistanceMeters
			durationSum = durationSum + statistics.TimeDuration
		}
		path[nextNode] = statistics
		departure = MinutesToCustomTime(CustomTimeToMinutes(departure) + int(statistics.TimeDuration.Minutes()))
		initNode = nextNode


	}

	return path, distanceSum, durationSum, departure
}

func MinutesToCustomTime(minutes int) string {
	hours := minutes/60
	minutesInHour := minutes%60
	return fmt.Sprintf("%02d", hours) + ":" + fmt.Sprintf("%02d", minutesInHour)
}

func OpeningHoursToMinutes(openingHours string) (int, int) {
	openingHoursStr := strings.Split(openingHours, "-")[0]
	closingHoursStr := strings.Split(openingHours, "-")[1]

	return CustomTimeToMinutes(openingHoursStr), CustomTimeToMinutes(closingHoursStr)
}

func CustomTimeToMinutes(customTime string) int {
	hourStr := strings.Split(customTime, ":")[0]
	minuteStr := strings.Split(customTime, ":")[1]

	hours, _ := strconv.Atoi(hourStr)
	minutes, _ := strconv.Atoi(minuteStr)

	return minutes + hours*60
}

func chooseTransportBetweenPlaces(place1 Place, place2 Place) string {
	return "driving"
}


var MOCK_MAP = map[Place]map[Place]RouteStatistics{
	{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}: {
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:237.5, Distance:1140, Time:832000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:1045.5416666666667, Distance:1091, Time:809000000000},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:99.375, Distance:477, Time:349000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:102.49999999999999, Distance:410, Time:307000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:171.33333333333334, Distance:514, Time:387000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:1236.25, Distance:1290, Time:938000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:1089.625, Distance:1137, Time:832000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:77.87499999999993, Distance:623, Time:451000000000},
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:598, Distance:624, Time:452000000000},
	},
	{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}: {
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:701.5, Distance:732, Time:596000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:246.99999999999997, Distance:988, Time:799000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:305.6666666666667, Distance:917, Time:738000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:715.875, Distance:747, Time:709000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:1006.25, Distance:1050, Time:836000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:1480.625, Distance:1545, Time:1187000000000},
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:129.79166666666666, Distance:623, Time:525000000000},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:121.875, Distance:585, Time:493000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:179.58333333333334, Distance:862, Time:798000000000},
	},
	{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}: {
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:237.5, Distance:1140, Time:864000000000},
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:1246.7916666666667, Distance:1301, Time:982000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:387.49999999999994, Distance:1550, Time:1172000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:175.375, Distance:183, Time:132000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:2025.9166666666667, Distance:2114, Time:1572000000000},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:240.41666666666669, Distance:1154, Time:879000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:108.49999999999997, Distance:868, Time:775000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:334, Distance:1002, Time:787000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:448.5, Distance:468, Time:358000000000},
	},
	{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}: {
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:230.41666666666669, Distance:1106, Time:847000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:131.25, Distance:1050, Time:786000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:97.5, Distance:468, Time:350000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:1978.9583333333333, Distance:2065, Time:1541000000000},
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:227.29166666666669, Distance:1091, Time:832000000000},
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:1199.8333333333333, Distance:1252, Time:950000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:375.24999999999983, Distance:1501, Time:1140000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:310.6666666666667, Distance:932, Time:727000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:592.25, Distance:618, Time:455000000000},
	},
	{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}: {
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:195, Distance:936, Time:686000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:196.74999999999997, Distance:787, Time:582000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:482.6666666666667, Distance:1448, Time:1078000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:1956.9166666666667, Distance:2042, Time:1510000000000},
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:236.875, Distance:1137, Time:840000000000},
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:756.125, Distance:789, Time:577000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:435.625, Distance:2091, Time:1533000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:2147.625, Distance:2241, Time:1639000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:190.125, Distance:1521, Time:1104000000000},
	},
	{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:    {Wight:99.375, Distance:477, Time:350000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                           {Wight:330.6666666666667, Distance:992, Time:737000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                        {Wight:240.41666666666669, Distance:1154, Time:848000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                   {Wight:920, Distance:960, Time:694000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:               {Wight:73.125, Distance:585, Time:419000000000},
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                  {Wight:140.875, Distance:147, Time:103000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:         {Wight:100.74999999999997, Distance:403, Time:306000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                            {Wight:1250.625, Distance:1305, Time:953000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}: {Wight:1059.9166666666667, Distance:1106, Time:824000000000},
	},
	{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}: {
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:123, Distance:492, Time:373000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:1390.5416666666667, Distance:1451, Time:1062000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:1199.8333333333333, Distance:1252, Time:933000000000},
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:130, Distance:624, Time:459000000000},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:30.625, Distance:147, Time:109000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:271.0416666666667, Distance:1301, Time:957000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:779.125, Distance:813, Time:591000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:91.5, Distance:732, Time:528000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:379.3333333333333, Distance:1138, Time:846000000000},
	},
	{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}: {
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:85.41666666666667, Distance:410, Time:298000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:123.49999999999997, Distance:988, Time:714000000000},
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:471.5, Distance:492, Time:356000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:224, Distance:672, Time:505000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:322.9166666666667, Distance:1550, Time:1130000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:1629.1666666666667, Distance:1700, Time:1236000000000},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:83.95833333333334, Distance:403, Time:295000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:1438.4583333333333, Distance:1501, Time:1106000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:754.2083333333334, Distance:787, Time:566000000000},
	},
	{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}: {
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:1090.5833333333333, Distance:1138, Time:828000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:168, Distance:672, Time:501000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:1386.7083333333333, Distance:1447, Time:1058000000000},
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:107.08333333333334, Distance:514, Time:376000000000},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:206.66666666666669, Distance:992, Time:725000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:114.62499999999993, Distance:917, Time:653000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:208.75, Distance:1002, Time:747000000000},
		{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}:                             {Wight:1104, Distance:1152, Time:852000000000},
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:893.1666666666666, Distance:932, Time:694000000000},
	},
	{Name:"Muzeum Wojska Polskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcRGDePQvvQqow"}: {
		{Name:"Muzeum Warszawskiego Przedsiębiorstwa Geodezyjnego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJ7-9EM_fMHkcRVh9fPLxfJW8"}:  {Wight:592.25, Distance:618, Time:474000000000},
		{Name:"Muzeum Teatralne", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJbQJ-qWbMHkcRrrYzTC9PLNw"}:                                    {Wight:2169.6666666666665, Distance:2264, Time:1689000000000},
		{Name:"Centrum Pieniądza NBP im. Sławomira S. Skrzypka", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJV9L_QvXMHkcR4fE_qS-WapY"}:     {Wight:268.75, Distance:1290, Time:980000000000},
		{Name:"Biuro Wystaw / Fundacja Polskiej Sztuki Nowoczesnej", OpeningHours:"11:00-18:00", Time:"1h", PlaceID:"ChIJj1ffSl7MHkcRBfCWp3BZgGc"}: {Wight:271.875, Distance:1305, Time:995000000000},
		{Name:"Muzeum Fryderyka Chopina w Warszawie", OpeningHours:"11:00-20:00", Time:"1h", PlaceID:"ChIJKbVhrFjMHkcRdbMJyIdXC34"}:                {Wight:94.125, Distance:753, Time:696000000000},
		{Name:"Muzeum Uniwersytetu Warszawskiego", OpeningHours:"00:00-24:00", Time:"1h", PlaceID:"ChIJnT24I17MHkcR8TirCHITPHM"}:                   {Wight:1390.5416666666667, Distance:1451, Time:1098000000000},
		{Name:"Państwowe Muzeum Etnograficzne w Warszawie", OpeningHours:"10:00-17:00", Time:"1h", PlaceID:"ChIJPcztF2DMHkcRNTXS6e61rqo"}:          {Wight:424.99999999999994, Distance:1700, Time:1288000000000},
		{Name:"Muzeum Historii Akimosik", OpeningHours:"10:00-15:00", Time:"1h", PlaceID:"ChIJZ1x3xfTMHkcRAimNnQ9D9os"}:                            {Wight:384, Distance:1152, Time:903000000000},
		{Name:"Muzeum Narodowe w Warszawie", OpeningHours:"10:00-18:00", Time:"1h", PlaceID:"ChIJK6UEcvfMHkcR0iGDJoZsIgc"}:                         {Wight:38.125, Distance:183, Time:143000000000},
	},
}