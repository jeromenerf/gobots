package angle

import (
	"math"
)

func Angle(x, y float64) float64 {
	angle := math.Atan2(x, y) / (math.Pi * 2) * 360
	direction := math.Mod((360 + 180 + angle), 360)
	return direction
}
