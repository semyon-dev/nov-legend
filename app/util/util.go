package util

import (
	"math"
)

func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	radlat1 := math.Pi * lat1 / 180
	radlat2 := math.Pi * lat2 / 180

	theta := lng1 - lng2
	radtheta := math.Pi * theta / 180

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	dist = dist * 1.609344 // kilometers

	return dist
}

// usage:

//winnipeg := coordinate{49.895077, -97.138451}
//regina := coordinate{50.445210, -104.618896}
//
//fmt.Println(distance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, "N"))
