package talibcdl

import (
	"testing"
)

func TestThreeInsideUp(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 175, 155, 185},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 170, 145, 146},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 100, 150, 180},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 140, 140},
	}
	cs := ThreeInside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3INSIDE(testOpen,testHigh,testLow,testClose)")
}

func TestThreeInsideDown(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 175, 155, 147},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 150, 146},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 170, 145, 80},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 140, 75},
	}
	cs := ThreeInside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3INSIDE(testOpen,testHigh,testLow,testClose)")
}
