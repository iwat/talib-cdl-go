package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - first candle: long white candle
 * - second candle: black real body
 * - gap between the first and the second candle's real bodies
 * - third candle: black candle that opens within the second real body and closes within the first real body
 * The meaning of "long" is specified with TA_SetCandleSettings
 * outInteger is negative (-1 to -100): two crows is always bearish;
 * the user should consider that two crows is significant when it appears in an uptrend, while this function
 * does not consider the trend
 */
func TestTwoCrows(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 215, 245, 235},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 110, 240, 230},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 200, 220, 150},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 215, 145},
	}
	cs := TwoCrows(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL2CROWS(testOpen,testHigh,testLow,testClose)")
}
