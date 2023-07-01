package model

import (
	"fmt"
	"math"
)

type Point struct {
	Latitude  int32
	Longitude int32
}

func (p *Point) Equals(q *Point) bool {
	if p.Latitude != q.Latitude ||
		p.Longitude != q.Longitude {

		return false
	}

	return true
}

func (p *Point) In(rect *Rectangle) bool {
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

func (p *Point) Serialize() string {
	return fmt.Sprintf("%d %d", p.Latitude, p.Longitude)
}
