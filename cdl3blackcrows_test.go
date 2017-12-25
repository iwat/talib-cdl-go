package talibcdl

import (
	"testing"
)

func TestThreeBlackCrows(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 165, 205, 170, 155},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 200, 175, 150},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 160, 150, 125, 100},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 149, 124, 105},
	}
	cs := ThreeBlackCrows(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3BLACKCROWS(testOpen,testHigh,testLow,testClose)")
}
