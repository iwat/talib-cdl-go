package talibcdl

import (
	"log"
)

// StickSandwich implements ta-lib function TA_CDLBREAKAWAY.
//
//         |       Up:              38%
//       | ▓       Down:            62%
//     | ░ ▓       Common Rank:     F-
//     ▓ ░ ▓       Efficiency Rank: B
//     ▓ ░ ▓       Source:          feedroll.com
//     ▓ | ▓
//         |
func StickSandwich(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingEqual.avgPeriod + 2

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
		equalPeriodTotal += es.rangeOf(settingEqual, i-2)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: black candle
	 * - second candle: white candle that trades only above the prior close (low > prior close)
	 * - third candle: black candle with the close equal to the first candle's close
	 * The meaning of "equal" is specified with TA_SetCandleSettings
	 * outInteger is always positive (1 to 100): stick sandwich is always bullish;
	 * the user should consider that stick sandwich is significant when comath.Ming in a downtrend,
	 * while this function does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-2).isBlack() && // first black
			es.candleColor(i-1).isWhite() && // second white
			es.candleColor(i).isBlack() && // third black
			es.Low(i-1) > es.Close(i-2) && // 2nd low > prior close
			es.Close(i) <= es.Close(i-2)+es.average(settingEqual, equalPeriodTotal, i-2) && // 1st and 3rd same close
			es.Close(i) >= es.Close(i-2)-es.average(settingEqual, equalPeriodTotal, i-2) {
			outInteger[i] = 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		equalPeriodTotal += es.rangeOf(settingEqual, i-2) - es.rangeOf(settingEqual, equalTrailingIdx-2)
		equalTrailingIdx++
	}

	return outInteger
}
