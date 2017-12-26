package talibcdl

import (
	"testing"
)

func TestBreakAway(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 205, 105, 104, 103, 108},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 200, 100, 99, 98, 95},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 105, 90, 89, 88, 103},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 85, 84, 83, 92},
	}
	cs := BreakAway(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLBREAKAWAY(testOpen,testHigh,testLow,testClose)")
}
