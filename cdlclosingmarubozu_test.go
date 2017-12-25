package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - long white (black) real body
 * - no or very short upper (lower) shadow
 * The meaning of "long" and "very short" is specified with TA_SetCandleSettings
 * outInteger is positive (1 to 100) when white (bullish), negative (-1 to -100) when black (bearish)
 */
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
