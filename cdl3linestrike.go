package talibcdl

import (
	"log"
	"math"
)

// ThreeLineStrike implements ta-lib function TA_CDL3LINESTRIKE.
//
//         | |
//       | ░ ▓
//     | ░ ░ ▓
//     ░ ░ | ▓
//     ░ |   ▓
//     |     ▓
//           |
func ThreeLineStrike(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingNear.avgPeriod + 3

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	nearTrailingIdx := startIdx - settingNear.avgPeriod

	nearPeriodTotal := [4]float64{}
	for i := nearTrailingIdx; i < startIdx; i++ {
		nearPeriodTotal[3] += es.rangeOf(settingNear, i-3)
		nearPeriodTotal[2] += es.rangeOf(settingNear, i-2)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - three white soldiers (three black crows): three white (black) candlesticks with consecutively higher (lower) closes,
	 * each opening within or near the previous real body
	 * - fourth candle: black (white) candle that opens above (below) prior candle's close and closes below (above)
	 * the first candle's open
	 * The meaning of "near" is specified with TA_SetCandleSettings;
	 * outInteger is positive (1 to 100) when bullish or negative (-1 to -100) when bearish;
	 * the user should consider that 3-line strike is significant when it appears in a trend in the same direction of
	 * the first three candles, while this function does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-3) == es.candleColor(i-2) && // three with same color
			es.candleColor(i-2) == es.candleColor(i-1) &&
			es.candleColor(i) == -es.candleColor(i-1) && // 4th opposite color
			// 2nd opens within/near 1st rb
			es.Open(i-2) >= math.Min(es.Open(i-3), es.Close(i-3))-es.average(settingNear, nearPeriodTotal[3], i-3) &&
			es.Open(i-2) <= math.Max(es.Open(i-3), es.Close(i-3))+es.average(settingNear, nearPeriodTotal[3], i-3) &&
			// 3rd opens within/near 2nd rb
			es.Open(i-1) >= math.Min(es.Open(i-2), es.Close(i-2))-es.average(settingNear, nearPeriodTotal[2], i-2) &&
			es.Open(i-1) <= math.Max(es.Open(i-2), es.Close(i-2))+es.average(settingNear, nearPeriodTotal[2], i-2) &&
			(( // if three white
			es.candleColor(i-1) == 1 &&
				es.Close(i-1) > es.Close(i-2) && es.Close(i-2) > es.Close(i-3) && // consecutive higher closes
				es.Open(i) > es.Close(i-1) && // 4th opens above prior close
				es.Close(i) < es.Open(i-3)) || // 4th closes below 1st open
				( // if three black
				es.candleColor(i-1) == -1 &&
					es.Close(i-1) < es.Close(i-2) && es.Close(i-2) < es.Close(i-3) && // consecutive lower closes
					es.Open(i) < es.Close(i-1) && // 4th opens below prior close
					es.Close(i) > es.Open(i-3))) { // 4th closes above 1st open
			outInteger[i] = int(es.candleColor(i-1)) * 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		for totIdx := 3; totIdx >= 2; totIdx-- {
			nearPeriodTotal[totIdx] += es.rangeOf(settingNear, i-totIdx) - es.rangeOf(settingNear, nearTrailingIdx-totIdx)
		}
		nearTrailingIdx++
	}

	return outInteger
}
