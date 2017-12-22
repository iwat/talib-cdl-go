package talibcdl

import (
	"log"
	"math"
)

// ThreeInside implements ta-lib function TA_CDL3INSIDE.
func ThreeInside(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingBodyShort.avgPeriod + 2
	if settingBodyShort.avgPeriod < settingBodyLong.avgPeriod {
		startIdx = settingBodyLong.avgPeriod + 2
	}

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyShortPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - 2 - settingBodyLong.avgPeriod
	bodyShortTrailingIdx := startIdx - 1 - settingBodyShort.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx-2; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i)
	}
	for i := bodyShortTrailingIdx; i < startIdx-1; i++ {
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: long white (black) real body
	 * - second candle: short real body totally engulfed by the first
	 * - third candle: black (white) candle that closes lower (higher) than the first candle's open
	 * The meaning of "short" and "long" is specified with TA_SetCandleSettings
	 * outInteger is positive (1 to 100) for the three inside up or negative (-1 to -100) for the three inside down;
	 * the user should consider that a three inside up is significant when it appears in a downtrend and a three inside
	 * down is significant when it appears in an uptrend, while this function does not consider the trend
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.realBody(i-2) > es.average(settingBodyLong, bodyLongPeriodTotal, i-2) && // 1st: long
			es.realBody(i-1) <= es.average(settingBodyShort, bodyShortPeriodTotal, i-1) && // 2nd: short
			math.Max(es.Close(i-1), es.Open(i-1)) < math.Max(es.Close(i-2), es.Open(i-2)) && // engulfed by 1st
			math.Min(es.Close(i-1), es.Open(i-1)) > math.Min(es.Close(i-2), es.Open(i-2)) &&
			((es.candleColor(i-2).IsWhite() && es.candleColor(i).IsBlack() && es.Close(i) < es.Open(i-2)) || // 3rd: opposite to 1st
				(es.candleColor(i-2).IsBlack() && es.candleColor(i).IsWhite() && es.Close(i) > es.Open(i-2))) { // and closing out
			outInteger[i] = -int(es.candleColor(i-2)) * 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-2) - es.rangeOf(settingBodyLong, bodyLongTrailingIdx)
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i-1) - es.rangeOf(settingBodyShort, bodyShortTrailingIdx)
		bodyLongTrailingIdx++
		bodyShortTrailingIdx++
	}

	return outInteger
}
