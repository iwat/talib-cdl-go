package talibcdl

import (
	"testing"
)

func TestDojiStar(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 165, 13},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 160, 12},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 110, 11},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 10},
	}
	cs := DojiStar(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLDOJISTAR(testOpen,testHigh,testLow,testClose)")
}
