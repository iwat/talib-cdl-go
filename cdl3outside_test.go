package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - first: black (white) real body
 * - second: white (black) real body that engulfs the prior real body
 * - third: candle that closes higher (lower) than the second candle
 * outInteger is positive (1 to 100) for the three outside up or negative (-1 to -100) for the three outside down;
 * the user should consider that a three outside up must appear in a downtrend and three outside down must appear
 * in an uptrend, while this function does not consider it
 */
func TestThreeOutsideUp(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 155, 165, 195},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 150, 100, 150},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 110, 160, 190},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 95, 145},
	}
	cs := ThreeOutside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3OUTSIDE(testOpen,testHigh,testLow,testClose)")
}

func TestThreeOutsideDown(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 155, 165, 115},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 110, 160, 110},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 150, 100, 50},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 95, 45},
	}
	cs := ThreeOutside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3OUTSIDE(testOpen,testHigh,testLow,testClose)")
}
