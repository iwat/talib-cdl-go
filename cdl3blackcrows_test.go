package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - three consecutive and declining black candlesticks
 * - each candle must have no or very short lower shadow
 * - each candle after the first must open within the prior candle's real body
 * - the first candle's close should be under the prior white candle's high
 * The meaning of "very short" is specified with TA_SetCandleSettings
 * outInteger is negative (-1 to -100): three black crows is always bearish;
 * the user should consider that 3 black crows is significant when it appears after a mature advance or at high levels,
 * while this function does not consider it
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
