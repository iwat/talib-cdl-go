package talibcdl

import (
	"log"
)

// ConcealBabySwall implements ta-lib function TA_CDL2CROWS.
//
//     ▓         Up:              25%
//     ▓         Down:            75%
//     ▓ ▓   ▓   Common Rank:     J-
//       ▓ | ▓   Efficiency Rank: J-
//         ▓ ▓   Source:          feedroll.com
func ConcealBabySwall(series Series) []int {
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
	shadowVeryShortPeriodTotal := [4]float64{}

	shadowVeryShortTrailingIdx := startIdx - settingShadowVeryShort.avgPeriod

	for i := shadowVeryShortTrailingIdx; i < startIdx; i++ {
		shadowVeryShortPeriodTotal[3] += es.rangeOf(settingShadowVeryShort, i-3)
		shadowVeryShortPeriodTotal[2] += es.rangeOf(settingShadowVeryShort, i-2)
		shadowVeryShortPeriodTotal[1] += es.rangeOf(settingShadowVeryShort, i-1)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: black marubozu (very short shadows)
	 * - second candle: black marubozu (very short shadows)
	 * - third candle: black candle that opens gapping down but has an upper shadow that extends into the prior body
	 * - fourth candle: black candle that completely engulfs the third candle, including the shadows
	 * The meanings of "very short shadow" are specified with TA_SetCandleSettings;
	 * outInteger is positive (1 to 100): concealing baby swallow is always bullish;
	 * the user should consider that concealing baby swallow is significant when it appears in downtrend, while
	 * this function does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-3).isBlack() && // 1st black
			es.candleColor(i-2).isBlack() && // 2nd black
			es.candleColor(i-1).isBlack() && // 3rd black
			es.candleColor(i).isBlack() && // 4th black
			// 1st: marubozu
			es.lowerShadow(i-3) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[3], i-3) &&
			es.upperShadow(i-3) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[3], i-3) &&
			// 2nd: marubozu
			es.lowerShadow(i-2) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[2], i-2) &&
			es.upperShadow(i-2) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[2], i-2) &&
			es.realBodyGapDown(i-1, i-2) && // 3rd: opens gapping down
			//      and HAS an upper shadow
			es.upperShadow(i-1) > es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[1], i-1) &&
			es.High(i-1) > es.Close(i-2) && //      that extends into the prior body
			es.High(i) > es.High(i-1) && es.Low(i) < es.Low(i-1) { // 4th: engulfs the 3rd including the shadows
			outInteger[i] = 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		for totIdx := 3; totIdx >= 1; totIdx-- {
			shadowVeryShortPeriodTotal[totIdx] += es.rangeOf(settingShadowVeryShort, i-totIdx) -
				es.rangeOf(settingShadowVeryShort, shadowVeryShortTrailingIdx-totIdx)
		}
		shadowVeryShortTrailingIdx++
	}

	return outInteger
}
