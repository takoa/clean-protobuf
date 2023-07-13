package model

import (
	"fmt"
	"math"
)

type Point struct {
	Latitude  int32
	Longitude int32
}

func (p Point) Equals(q Point) bool {
	if p.Latitude != q.Latitude ||
		p.Longitude != q.Longitude {

		return false
	}

	return true
}

func (p Point) Distance(p2 Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

func (p Point) In(rect Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if left <= float64(p.Longitude) && float64(p.Longitude) <= right &&
		bottom <= float64(p.Latitude) && float64(p.Latitude) <= top {
		return true
	}

	return false
}

func (p Point) Serialize() string {
	return fmt.Sprintf("%d %d", p.Latitude, p.Longitude)
}
