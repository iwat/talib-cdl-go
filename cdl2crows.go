package talibcdl

import (
	"log"
)

// TwoCrows implements ta-lib function TA_CDL2CROWS.
func TwoCrows(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingBodyLong.avgPeriod + 2

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - 2 - settingBodyLong.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx-2; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: long white candle
	 * - second candle: black real body
	 * - gap between the first and the second candle's real bodies
	 * - third candle: black candle that opens within the second real body and closes within the first real body
	 * The meaning of "long" is specified with TA_SetCandleSettings
	 * outInteger is negative (-1 to -100): two crows is always bearish;
	 * the user should consider that two crows is significant when it appears in an uptrend, while this function
	 * does not consider the trend
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-2).IsWhite() &&
			es.realBody(i-2) > es.average(settingBodyLong, bodyLongPeriodTotal, i-2) && // long
			es.candleColor(i-1).IsBlack() &&
			es.realBodyGapUp(i-1, i-2) && // gapping up
			es.candleColor(i).IsBlack() &&
			es.Open(i) < es.Open(i-1) && es.Open(i) > es.Close(i-1) && // opening within 2nd rb
			es.Close(i) > es.Open(i-2) && es.Close(i) < es.Close(i-2) { // closing within 1st rb
			outInteger[i] = -100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-2) - es.rangeOf(settingBodyLong, bodyLongTrailingIdx)
		bodyLongTrailingIdx++
	}

	return outInteger
}
