package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - first candle: long white (black) real body
 * - second candle: short real body totally engulfed by the first
 * - third candle: black (white) candle that closes lower (higher) than the first candle's open
 * The meaning of "short" and "long" is specified with TA_SetCandleSettings
 * outInteger is positive (1 to 100) for the three inside up or negative (-1 to -100) for the three inside down;
 * the user should consider that a three inside up is significant when it appears in a downtrend and a three inside
 * down is significant when it appears in an uptrend, while this function does not consider the trend
 */
func TestThreeInsideUp(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 175, 155, 185},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 170, 145, 146},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 150, 180},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 140, 140},
	}
	cs := ThreeInside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3INSIDE(testOpen,testHigh,testLow,testClose)")
}

func TestThreeInsideDown(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 175, 155, 147},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 150, 146},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 170, 145, 80},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 140, 75},
	}
	cs := ThreeInside(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3INSIDE(testOpen,testHigh,testLow,testClose)")
}
