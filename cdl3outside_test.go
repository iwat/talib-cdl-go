package talibcdl

import (
	"testing"
)

func TestThreeOutsideUp(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 155, 165, 195},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 150, 100, 150},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 110, 160, 190},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 95, 145},
	}
	cs := ThreeOutside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3OUTSIDE(testOpen,testHigh,testLow,testClose)")
}

func TestThreeOutsideDown(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 155, 165, 115},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 110, 160, 110},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 150, 100, 50},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 95, 45},
	}
	cs := ThreeOutside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3OUTSIDE(testOpen,testHigh,testLow,testClose)")
}
