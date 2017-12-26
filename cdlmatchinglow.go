package talibcdl

import (
	"log"
)

// MatchingLow implements ta-lib function TA_CDLMATCHINGLOW.
//
//     |     Up:              39%
//     ▓ |   Down:            61%
//     ▓ ▓   Common Rank:     F-
//     ▓ ▓   Efficiency Rank: A-
//     ▓ ▓   Source:          feedroll.com
//     | |
func MatchingLow(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingEqual.avgPeriod + 1

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	equalPeriodTotal := 0.0
	equalTrailingIdx := startIdx - settingEqual.avgPeriod

	for i := equalTrailingIdx; i < startIdx; i++ {
		equalPeriodTotal += es.rangeOf(settingEqual, i-1)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: black candle
	 * - second candle: black candle with the close equal to the previous close
	 * The meaning of "equal" is specified with TA_SetCandleSettings
	 * outInteger is always positive (1 to 100): matching low is always bullish;
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-1) == -1 && // first black
			es.candleColor(i) == -1 && // second black
			es.Close(i) <= es.Close(i-1)+es.average(settingEqual, equalPeriodTotal, i-1) && // 1st and 2nd same close
			es.Close(i) >= es.Close(i-1)-es.average(settingEqual, equalPeriodTotal, i-1) {
			outInteger[i] = 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		equalPeriodTotal += es.rangeOf(settingEqual, i-1) - es.rangeOf(settingEqual, equalTrailingIdx-1)
		equalTrailingIdx++
	}

	return outInteger
}
