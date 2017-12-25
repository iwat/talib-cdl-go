package talibcdl

import (
	"log"
)

// BeltHold implements ta-lib function TA_CDLBELTHOLD.
//
//     .
//     ░
//     ░
//     ░
func BeltHold(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := intMax(settingBodyLong.avgPeriod, settingShadowVeryShort.avgPeriod)

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx+1)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - settingBodyLong.avgPeriod
	shadowVeryShortPeriodTotal := 0.0
	shadowVeryShortTrailingIdx := startIdx - settingShadowVeryShort.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i)
	}
	for i := shadowVeryShortTrailingIdx; i < startIdx; i++ {
		shadowVeryShortPeriodTotal += es.rangeOf(settingShadowVeryShort, i)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - long white (black) real body
	 * - no or very short lower (upper) shadow
	 * The meaning of "long" and "very short" is specified with TA_SetCandleSettings
	 * outInteger is positive (1 to 100) when white (bullish), negative (-1 to -100) when black (bearish)
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.realBody(i) > es.average(settingBodyLong, bodyLongPeriodTotal, i) && // long body
			(( // white body and very short lower shadow
			es.candleColor(i).isWhite() &&
				es.lowerShadow(i) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal, i)) ||
				( // black body and very short upper shadow
				es.candleColor(i).isBlack() &&
					es.upperShadow(i) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal, i))) {
			outInteger[i] = int(es.candleColor(i)) * 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i) -
			es.rangeOf(settingBodyLong, bodyLongTrailingIdx)
		shadowVeryShortPeriodTotal += es.rangeOf(settingShadowVeryShort, i) -
			es.rangeOf(settingShadowVeryShort, shadowVeryShortTrailingIdx)
		bodyLongTrailingIdx++
		shadowVeryShortTrailingIdx++
	}

	return outInteger
}
