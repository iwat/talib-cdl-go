package talibcdl

import (
	"log"
)

// ThreeBlackCrows implements ta-lib function TA_CDL3BLACKCROWS.
//
//     ▓
//     ▓ ▓
//       ▓ ▓
//         ▓
func ThreeBlackCrows(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingShadowVeryShort.avgPeriod + 3

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	shadowVeryShortTrailingIdx := startIdx - settingShadowVeryShort.avgPeriod

	shadowVeryShortPeriodTotal := [3]float64{}
	for i := shadowVeryShortTrailingIdx; i < startIdx; i++ {
		shadowVeryShortPeriodTotal[2] += es.rangeOf(settingShadowVeryShort, i-2)
		shadowVeryShortPeriodTotal[1] += es.rangeOf(settingShadowVeryShort, i-1)
		shadowVeryShortPeriodTotal[0] += es.rangeOf(settingShadowVeryShort, i)
	}

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
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-3).isWhite() &&
			es.candleColor(i-2).isBlack() &&
			es.lowerShadow(i-2) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[2], i-2) &&
			// very short lower shadow
			es.candleColor(i-1).isBlack() &&
			es.lowerShadow(i-1) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[1], i-1) &&
			// very short lower shadow
			es.candleColor(i).isBlack() &&
			es.lowerShadow(i) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[0], i) &&
			// very short lower shadow
			es.Open(i-1) < es.Open(i-2) && es.Open(i-1) > es.Close(i-2) && // 2nd black opens within 1st black's rb
			es.Open(i) < es.Open(i-1) && es.Open(i) > es.Close(i-1) && // 3rd black opens within 2nd black's rb
			es.High(i-3) > es.Close(i-2) && // 1st black closes under prior candle's high
			es.Close(i-2) > es.Close(i-1) && // three declining
			es.Close(i-1) > es.Close(i) { // three declining
			outInteger[i] = -100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		for totIdx := 2; totIdx > 0; totIdx-- {
			shadowVeryShortPeriodTotal[totIdx] += es.rangeOf(settingShadowVeryShort, i-totIdx) -
				es.rangeOf(settingShadowVeryShort, shadowVeryShortTrailingIdx-totIdx)
		}
		shadowVeryShortTrailingIdx++
	}

	return outInteger
}
