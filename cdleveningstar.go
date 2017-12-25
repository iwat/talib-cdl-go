package talibcdl

import (
	"log"
)

// EveningStar implements ta-lib function TA_CDLEVENINGSTAR.
//
//       |     Up:              28%
//     | ░ |   Down:            72%
//     ░ | ▓   Common Rank:     H+
//     ░   ▓   Efficiency Rank: A
//     |   ▓   Source:          feedroll.com
//         |
func EveningStar(series Series, penetration float64) []int {
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
	startIdx := intMax(settingBodyShort.avgPeriod, settingBodyLong.avgPeriod) + 2

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyShortPeriodTotal := 0.0
	bodyShortPeriodTotal2 := 0.0
	bodyLongTrailingIdx := startIdx - 2 - settingBodyLong.avgPeriod
	bodyShortTrailingIdx := startIdx - 1 - settingBodyShort.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx-2; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i)
	}
	for i := bodyShortTrailingIdx; i < startIdx-1; i++ {
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i)
		bodyShortPeriodTotal2 += es.rangeOf(settingBodyShort, i+1)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: long white real body
	 * - second candle: star (short real body gapping up)
	 * - third candle: black real body that moves well within the first candle's real body
	 * The meaning of "short" and "long" is specified with TA_SetCandleSettings
	 * The meaning of "moves well within" is specified with optInPenetration and "moves" should mean the real body should
	 * not be short ("short" is specified with TA_SetCandleSettings) - Greg Morris wants it to be long, someone else want
	 * it to be relatively long
	 * outInteger is negative (-1 to -100): evening star is always bearish;
	 * the user should consider that an evening star is significant when it appears in an uptrend,
	 * while this function does not consider the trend
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.realBody(i-2) > es.average(settingBodyLong, bodyLongPeriodTotal, i-2) && // 1st: long
			es.candleColor(i-2) == 1 && // white
			es.realBody(i-1) <= es.average(settingBodyShort, bodyShortPeriodTotal, i-1) && // 2nd: short
			es.realBodyGapUp(i-1, i-2) && // gapping up
			es.realBody(i) > es.average(settingBodyShort, bodyShortPeriodTotal2, i) && // 3rd: longer than short
			es.candleColor(i) == -1 && // black real body
			es.Close(i) < es.Close(i-2)-es.realBody(i-2)*penetration { // closing well within 1st rb
			outInteger[i] = -100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-2) -
			es.rangeOf(settingBodyLong, bodyLongTrailingIdx)
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i-1) -
			es.rangeOf(settingBodyShort, bodyShortTrailingIdx)
		bodyShortPeriodTotal2 += es.rangeOf(settingBodyShort, i) -
			es.rangeOf(settingBodyShort, bodyShortTrailingIdx+1)
		bodyLongTrailingIdx++
		bodyShortTrailingIdx++
	}

	return outInteger
}
