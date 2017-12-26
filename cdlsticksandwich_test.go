package talibcdl

import (
	"testing"
)

func TestStickSandwich(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 205, 230, 240},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 200, 110, 230},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 105, 220, 105},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 104, 106, 104},
	}
	cs := StickSandwich(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLSTICKSANDWICH(testOpen,testHigh,testLow,testClose)")
}
