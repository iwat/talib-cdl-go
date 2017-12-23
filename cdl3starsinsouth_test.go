package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - first candle: long black candle with long lower shadow
 * - second candle: smaller black candle that opens higher than prior close but within prior candle's range
 *   and trades lower than prior close but not lower than prior low and closes off of its low (it has a shadow)
 * - third candle: small black marubozu (or candle with very short shadows) engulfed by prior candle's range
 * The meanings of "long body", "short body", "very short shadow" are specified with TA_SetCandleSettings;
 * outInteger is positive (1 to 100): 3 stars in the south is always bullish;
 * the user should consider that 3 stars in the south is significant when it appears in downtrend, while this function
 * does not consider it
 */
func TestThreeStarsInSouth(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 205, 155, 144.1},
		Opens:  []float64{150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 200, 150, 144},
		Closes: []float64{150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 120, 130, 135},
		Lows:   []float64{145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 10, 109, 134.9},
	}
	cs := ThreeStarsInSouth(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3STARSINSOUTH(testOpen,testHigh,testLow,testClose)")
}
