package talibcdl

import (
	"testing"
)

func TestThreeStarsInSouth(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 205, 155, 144.1},
		Opens:  []float64{150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 200, 150, 144},
		Closes: []float64{150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 120, 130, 135},
		Lows:   []float64{145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 10, 109, 134.9},
	}
	cs := ThreeStarsInSouth(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3STARSINSOUTH(testOpen,testHigh,testLow,testClose)")
}
