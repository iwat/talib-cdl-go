package talibcdl

import (
	"testing"
)

func TestTwoCrows(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 215, 245, 235},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 110, 240, 230},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 200, 220, 150},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 215, 145},
	}
	cs := TwoCrows(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL2CROWS(testOpen,testHigh,testLow,testClose)")
}
