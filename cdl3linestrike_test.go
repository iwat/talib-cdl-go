package talibcdl

import (
	"testing"
)

func TestThreeLineStrikeUp(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 155, 165, 175, 185},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 110, 120, 180},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 150, 160, 170, 90},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 105, 115, 85},
	}
	cs := ThreeLineStrike(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3LINESTRIKE(testOpen,testHigh,testLow,testClose)")
}

func TestThreeLineStrikeDown(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 175, 165, 155, 185},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 170, 160, 150, 90},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 120, 110, 100, 180},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 115, 105, 95, 85},
	}
	cs := ThreeLineStrike(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3LINESTRIKE(testOpen,testHigh,testLow,testClose)")
}
