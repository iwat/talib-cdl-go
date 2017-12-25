package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - first candle: black marubozu (very short shadows)
 * - second candle: black marubozu (very short shadows)
 * - third candle: black candle that opens gapping down but has an upper shadow that extends into the prior body
 * - fourth candle: black candle that completely engulfs the third candle, including the shadows
 * The meanings of "very short shadow" are specified with TA_SetCandleSettings;
 * outInteger is positive (1 to 100): concealing baby swallow is always bullish;
 * the user should consider that concealing baby swallow is significant when it appears in downtrend, while
 * this function does not consider it
 */
func TestConcealBabySwall(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 201, 191, 150, 186},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 200, 190, 110, 185},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 130, 120, 100, 90},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 129, 119, 99, 89},
	}
	cs := ConcealBabySwall(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLCONCEALBABYSWALL(testOpen,testHigh,testLow,testClose)")
}
