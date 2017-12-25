package talibcdl

import (
	"testing"
)

func TestConcealBabySwall(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 201, 191, 150, 186},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 200, 190, 110, 185},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 130, 120, 100, 90},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 129, 119, 99, 89},
	}
	cs := ConcealBabySwall(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLCONCEALBABYSWALL(testOpen,testHigh,testLow,testClose)")
}
