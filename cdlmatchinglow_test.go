package talibcdl

import (
	"testing"
)

func TestMatchingLow(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 205, 185},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 200, 180},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 105, 105},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 100},
	}
	cs := MatchingLow(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLMATCHINGLOW(testOpen,testHigh,testLow,testClose)")
}
