package talibcdl

import (
	"testing"
)

func TestBeltHold(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 161},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 101},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 160},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100},
	}
	cs := BeltHold(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLBELTHOLD(testOpen,testHigh,testLow,testClose)")
}
