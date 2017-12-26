package talibcdl

import (
	"testing"
)

func TestClosingMarubozu(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 161, 200},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 161},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 160, 101},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 90, 102},
	}
	cs := ClosingMarubozu(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLCLOSINGMARUBOZU(testOpen,testHigh,testLow,testClose)")
}
