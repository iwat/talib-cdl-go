package talibcdl

import (
	"log"
)

// Piercing implements ta-lib function TA_CDLPIERCING.
//
//     |     Up:              64%
//     ▓ |   Down:            39%
//     ▓ ░   Common Rank:     D-
//     ▓ ░   Efficiency Rank: B+
//     | ░   Source:          feedroll.com
//       |
func Piercing(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingBodyLong.avgPeriod + 1

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := [2]float64{}
	bodyLongTrailingIdx := startIdx - settingBodyLong.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx; i++ {
		bodyLongPeriodTotal[1] += es.rangeOf(settingBodyLong, i-1)
		bodyLongPeriodTotal[0] += es.rangeOf(settingBodyLong, i)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: long black candle
	 * - second candle: long white candle with open below previous day low and close at least at 50% of previous day
	 * real body
	 * The meaning of "long" is specified with TA_SetCandleSettings
	 * outInteger is positive (1 to 100): piercing pattern is always bullish
	 * the user should consider that a piercing pattern is significant when it appears in a downtrend, while
	 * this function does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-1) == -1 && // 1st: black
			es.realBody(i-1) > es.average(settingBodyLong, bodyLongPeriodTotal[1], i-1) && // long
			es.candleColor(i) == 1 && // 2nd: white
			es.realBody(i) > es.average(settingBodyLong, bodyLongPeriodTotal[0], i) && // long
			es.Open(i) < es.Low(i-1) && // open below prior low
			es.Close(i) < es.Open(i-1) && // close within prior body
			es.Close(i) > es.Close(i-1)+es.realBody(i-1)*0.5 { // above midpoint
			outInteger[i] = 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		for totIdx := 1; totIdx >= 0; totIdx-- {
			bodyLongPeriodTotal[totIdx] += es.rangeOf(settingBodyLong, i-totIdx) -
				es.rangeOf(settingBodyLong, bodyLongTrailingIdx-totIdx)
		}
		bodyLongTrailingIdx++
	}

	return outInteger
}
