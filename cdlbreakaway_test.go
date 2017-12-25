package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - first candle: long black (white)
 * - second candle: black (white) day whose body gaps down (up)
 * - third candle: black or white day with lower (higher) high and lower (higher) low than prior candle's
 * - fourth candle: black (white) day with lower (higher) high and lower (higher) low than prior candle's
 * - fifth candle: white (black) day that closes inside the gap, erasing the prior 3 days
 * The meaning of "long" is specified with TA_SetCandleSettings
 * outInteger is positive (1 to 100) when bullish or negative (-1 to -100) when bearish;
 * the user should consider that breakaway is significant in a trend opposite to the last candle, while this
 * function does not consider it
 */
func TestBreakAway(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 205, 105, 104, 103, 108},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 200, 100, 99, 98, 95},
		Closes: []float64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 105, 90, 89, 88, 103},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 85, 84, 83, 92},
	}
	cs := BreakAway(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDLBREAKAWAY(testOpen,testHigh,testLow,testClose)")
}
