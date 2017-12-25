package talibcdl

import (
	"testing"
)

func TestEveningStar(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 155, 176, 155},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 170, 150},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 150, 171, 100},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 165, 105},
	}
	cs := EveningStar(d, DefaultFloat64)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLEVENINGSTAR(testOpen,testHigh,testLow,testClose)")
}
