package talibcdl

import (
	"log"
)

// AbandonedBaby implements ta-lib function TA_CDLABANDONEDBABY.
//
//       +
//         |
//     |   ▓
//     ░   ▓
//     ░   ▓
//     ░   |
//     |
func AbandonedBaby(series Series, penetration float64) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	if penetration == DefaultFloat64 {
		penetration = 3.000000e-1
	} else if penetration < 0.000000e+0 || penetration > 3.000000e+37 {
		log.Printf("penetration out of range")
		return outInteger
	}

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := intMax(
		intMax(settingBodyDoji.avgPeriod, settingBodyLong.avgPeriod),
		settingBodyShort.avgPeriod) + 2

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyDojiPeriodTotal := 0.0
	bodyShortPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - 2 - settingBodyLong.avgPeriod
	bodyDojiTrailingIdx := startIdx - 1 - settingBodyDoji.avgPeriod
	bodyShortTrailingIdx := startIdx - settingBodyShort.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx-2; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i)
	}
	for i := bodyDojiTrailingIdx; i < startIdx-1; i++ {
		bodyDojiPeriodTotal += es.rangeOf(settingBodyDoji, i)
	}
	for i := bodyShortTrailingIdx; i < startIdx; i++ {
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i)
	}

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
	for i := startIdx; i < es.Len(); i++ {
		if es.realBody(i-2) > es.average(settingBodyLong, bodyLongPeriodTotal, i-2) && // 1st: long
			es.realBody(i-1) <= es.average(settingBodyDoji, bodyDojiPeriodTotal, i-1) && // 2nd: doji
			es.realBody(i) > es.average(settingBodyShort, bodyShortPeriodTotal, i) && // 3rd: longer than short
			((es.candleColor(i-2) == 1 && // 1st white
				es.candleColor(i) == -1 && // 3rd black
				es.Close(i) < es.Close(i-2)-es.realBody(i-2)*penetration && // 3rd closes well within 1st rb
				es.isCandleGapUp(i-1, i-2) && // upside gap between 1st and 2nd
				es.isCandleGapDown(i, i-1)) || // downside gap between 2nd and 3rd
				(es.candleColor(i-2) == -1 && // 1st black
					es.candleColor(i) == 1 && // 3rd white
					es.Close(i) > es.Close(i-2)+es.realBody(i-2)*penetration && // 3rd closes well within 1st rb
					es.isCandleGapDown(i-1, i-2) && // downside gap between 1st and 2nd
					es.isCandleGapUp(i, i-1))) { // upside gap between 2nd and 3rd
			outInteger[i] = int(es.candleColor(i)) * 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-2) - es.rangeOf(settingBodyLong, bodyLongTrailingIdx)
		bodyDojiPeriodTotal += es.rangeOf(settingBodyDoji, i-1) - es.rangeOf(settingBodyDoji, bodyDojiTrailingIdx)
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i) - es.rangeOf(settingBodyShort, bodyShortTrailingIdx)
		bodyLongTrailingIdx++
		bodyDojiTrailingIdx++
		bodyShortTrailingIdx++
	}

	return outInteger
}
