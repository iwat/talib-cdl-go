package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - first candle: long white (black) real body
 * - second candle: doji
 * - third candle: black (white) real body that moves well within the first candle's real body
 * - upside (downside) gap between the first candle and the doji (the shadows of the two candles don't touch)
 * - downside (upside) gap between the doji and the third candle (the shadows of the two candles don't touch)
 * The meaning of "doji" and "long" is specified with TA_SetCandleSettings
 * The meaning of "moves well within" is specified with optInPenetration and "moves" should mean the real body should
 * not be short ("short" is specified with TA_SetCandleSettings) - Greg Morris wants it to be long, someone else want
 * it to be relatively long
 * outInteger is positive (1 to 100) when it's an abandoned baby bottom or negative (-1 to -100) when it's
 * an abandoned baby top; the user should consider that an abandoned baby is significant when it appears in
 * an uptrend or downtrend, while this function does not consider the trend
 */
func TestAbandonedBaby(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 155, 176, 155},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 170, 150},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 150, 170.1, 100},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 165, 105},
	}
	cs := AbandonedBaby(d, DefaultFloat64)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLABANDONEDBABY(testOpen,testHigh,testLow,testClose)")
}
