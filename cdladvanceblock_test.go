package talibcdl

import (
	"testing"
)

func TestAdvanceBlock(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 150.1, 190, 189},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 150, 161},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 150, 165, 170},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 99, 150, 160.9},
	}
	cs := AdvanceBlock(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLADVANCEBLOCK(testOpen,testHigh,testLow,testClose)")
}
