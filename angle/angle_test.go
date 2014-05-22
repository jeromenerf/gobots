package angle

import "testing"

var tests = [][]float64{
	[]float64{0, 0, 0},
	[]float64{0, 1, 180},
	[]float64{0, -1, 0},
	[]float64{1, 0, 270},
	[]float64{1, 1, 225},
	[]float64{1, -1, 315},
	[]float64{-1, 0, 90},
	[]float64{-1, 1, 135},
	[]float64{-1, -1, 45},
}

func TestAngle(t *testing.T) {
	for _, pair := range tests {
		v := Angle(pair[0], pair[1])
		if v != pair[2] {
			t.Error(
				"For", pair[:2],
				"expected", pair[2],
				"got", v,
			)
		}
	}
}
