package talibcdl

import (
	"testing"
)

/* Proceed with the calculation for the requested range.
 * Must have:
 * - three white soldiers (three black crows): three white (black) candlesticks with consecutively higher (lower) closes,
 * each opening within or near the previous real body
 * - fourth candle: black (white) candle that opens above (below) prior candle's close and closes below (above)
 * the first candle's open
 * The meaning of "near" is specified with TA_SetCandleSettings;
 * outInteger is positive (1 to 100) when bullish or negative (-1 to -100) when bearish;
 * the user should consider that 3-line strike is significant when it appears in a trend in the same direction of
 * the first three candles, while this function does not consider it
 */
func TestThreeLineStrikeUp(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 155, 165, 175, 185},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 100, 110, 120, 180},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 150, 160, 170, 90},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 95, 105, 115, 85},
	}
	cs := ThreeLineStrike(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3LINESTRIKE(testOpen,testHigh,testLow,testClose)")
}

func TestThreeLineStrikeDown(t *testing.T) {
	d := SimpleSeries{
		Highs:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 175, 165, 155, 185},
		Opens:  []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 170, 160, 150, 90},
		Closes: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 120, 110, 100, 180},
		Lows:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 115, 105, 95, 85},
	}
	cs := ThreeLineStrike(d)
	t.Log(cs)

	compareInts(t, d, cs, "result = talib.CDL3LINESTRIKE(testOpen,testHigh,testLow,testClose)")
}
