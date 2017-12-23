package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - three white candlesticks with consecutively higher closes
 * - Greg Morris wants them to be long, Steve Nison doesn't; anyway they should not be short
 * - each candle opens within or near the previous white real body
 * - each candle must have no or very short upper shadow
 * - to differentiate this pattern from advance block, each candle must not be far shorter than the prior candle
 * The meanings of "not short", "very short shadow", "far" and "near" are specified with TA_SetCandleSettings;
 * here the 3 candles must be not short, if you want them to be long use TA_SetCandleSettings on BodyShort;
 * outInteger is positive (1 to 100): advancing 3 white soldiers is always bullish;
 * the user should consider that 3 white soldiers is significant when it appears in downtrend, while this function
 * does not consider it
 */
func TestThreeWhiteSoldiers(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 130.1, 150.1, 170.1},
		Opens:  []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 120, 140},
		Closes: []float64{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 130, 150, 170},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 99.9, 119.9, 139.9},
	}
	cs := ThreeWhiteSoldiers(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3WHITESOLDIERS(testOpen,testHigh,testLow,testClose)")
}
