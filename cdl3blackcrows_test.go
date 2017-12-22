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
 * ThreeBlack is negative (-1 to -100): two crows is always bearish;
 * the user should consider that two crows is significant when it appears in an uptrend, while this function
 * does not consider the trend
 */
func TestThreeBlackCrows(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 165, 205, 170, 155},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 200, 175, 150},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 160, 150, 125, 100},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 105, 149.9, 124.9, 105},
	}
	cs := ThreeBlackCrows(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3BLACKCROWS(testOpen,testHigh,testLow,testClose)")
}
