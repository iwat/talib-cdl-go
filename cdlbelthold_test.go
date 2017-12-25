package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - long white (black) real body
 * - no or very short lower (upper) shadow
 * The meaning of "long" and "very short" is specified with TA_SetCandleSettings
 * outInteger is positive (1 to 100) when white (bullish), negative (-1 to -100) when black (bearish)
 */
func TestBeltHold(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 161},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 101},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 160},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100},
	}
	cs := BeltHold(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLBELTHOLD(testOpen,testHigh,testLow,testClose)")
}
