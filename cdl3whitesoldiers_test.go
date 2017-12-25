package talibcdl

import (
	"testing"
)

func TestThreeWhiteSoldiers(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 131, 151, 171},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 120, 140},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 130, 150, 170},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 99, 119, 139},
	}
	cs := ThreeWhiteSoldiers(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3WHITESOLDIERS(testOpen,testHigh,testLow,testClose)")
}
