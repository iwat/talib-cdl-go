package talibcdl

import (
	"log"
)

// ThreeStarsInSouth implements ta-lib function TA_CDL3STARSINSOUTH.
func ThreeStarsInSouth(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := intMax(
		intMax(settingShadowVeryShort.avgPeriod, settingShadowLong.avgPeriod),
		intMax(settingBodyLong.avgPeriod, settingBodyShort.avgPeriod)) + 2

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - settingBodyLong.avgPeriod
	shadowLongPeriodTotal := 0.0
	shadowLongTrailingIdx := startIdx - settingShadowLong.avgPeriod
	shadowVeryShortTrailingIdx := startIdx - settingShadowVeryShort.avgPeriod
	bodyShortPeriodTotal := 0.0
	bodyShortTrailingIdx := startIdx - settingBodyShort.avgPeriod

	shadowVeryShortPeriodTotal := [2]float64{}
	for i := bodyLongTrailingIdx; i < startIdx; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-2)
	}
	for i := shadowLongTrailingIdx; i < startIdx; i++ {
		shadowLongPeriodTotal += es.rangeOf(settingShadowLong, i-2)
	}
	for i := shadowVeryShortTrailingIdx; i < startIdx; i++ {
		shadowVeryShortPeriodTotal[1] += es.rangeOf(settingShadowVeryShort, i-1)
		shadowVeryShortPeriodTotal[0] += es.rangeOf(settingShadowVeryShort, i)
	}
	for i := bodyShortTrailingIdx; i < startIdx; i++ {
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: long black candle with long lower shadow
	 * - second candle: smaller black candle that opens higher than prior close but within prior candle's range
	 *   and trades lower than prior close but not lower than prior low and closes off of its low (it has a shadow)
	 * - third candle: small black marubozu (or candle with very short shadows) engulfed by prior candle's range
	 * The meanings of "long body", "short body", "very short shadow" are specified with TA_SetCandleSettings;
	 * outInteger is positive (1 to 100): 3 stars in the south is always bullish;
	 * the user should consider that 3 stars in the south is significant when it appears in downtrend, while this function
	 * does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-2).isBlack() && // 1st black
			es.candleColor(i-1).isBlack() && // 2nd black
			es.candleColor(i).isBlack() && // 3rd black
			// 1st: long
			es.realBody(i-2) > es.average(settingBodyLong, bodyLongPeriodTotal, i-2) &&
			// with long lower shadow
			es.lowerShadow(i-2) > es.average(settingShadowLong, shadowLongPeriodTotal, i-2) &&
			es.realBody(i-1) < es.realBody(i-2) && // 2nd: smaller candle
			es.Open(i-1) > es.Close(i-2) && es.Open(i-1) <= es.High(i-2) && // that opens higher but within 1st range
			es.Low(i-1) < es.Close(i-2) && // and trades lower than 1st close
			es.Low(i-1) >= es.Low(i-2) && // but not lower than 1st low
			// and has a lower shadow
			es.lowerShadow(i-1) > es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[1], i-1) &&
			// 3rd: small marubozu
			es.realBody(i) < es.average(settingBodyShort, bodyShortPeriodTotal, i) &&
			es.lowerShadow(i) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[0], i) &&
			es.upperShadow(i) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[0], i) &&
			es.Low(i) > es.Low(i-1) && es.High(i) < es.High(i-1) { // engulfed by prior candle's range
			outInteger[i] = 100
		}
	}

	return outInteger
}
