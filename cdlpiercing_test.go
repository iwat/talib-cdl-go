package talibcdl

import (
	"testing"
)

func TestPiercing(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 205, 195},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 200, 104},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 110, 185},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 100},
	}
	cs := Piercing(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLPIERCING(testOpen,testHigh,testLow,testClose)")
}
