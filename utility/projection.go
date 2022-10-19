package utility

import (
	"fmt"
	"math"
	"strconv"
)

func ConvertLatLonToXY(lat float64, lon float64) (x float64, y float64) {
	tempLat := fmt.Sprintf("%f", lat)
	tempLon := fmt.Sprintf("%f", lon)
	tLat := tempLat[5:]
	tLon := tempLon[5:]

	strLat := tLat[0:2] + "." + tLat[2:]
	strLon := tLon[0:2] + "." + tLon[2:]

	testLat, _ := strconv.ParseFloat(strLat, 64)
	testLon, _ := strconv.ParseFloat(strLon, 64)

	fmt.Println("lat: ", tLat, "lon: ", tLon)
	x = ((testLon + 180) * (1024 / 360))
	y = (((testLat * -1) + 90) * (1024 / 180))
	return x, y
}

func radian(d float64) float64 {
	return d * math.Pi / 180
}

func LonLatToPixelMercator(lon, lat, tile, zoom float64) (px, py float64) {
	east := radian(lon) * 6378137
	north := math.Log(math.Tan(radian((90+lat)/2))) * 6378137

	res := (2 * math.Pi * 6378137 / tile) / math.Pow(2, zoom)
	shift := math.Pi * 6378137

	return (east + shift) / res, (north + shift) / res
}
