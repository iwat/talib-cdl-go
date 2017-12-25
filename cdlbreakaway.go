package talibcdl

import (
	"log"
)

// BreakAway implements ta-lib function TA_CDLBREAKAWAY.
//
//     |           Up:              63%
//     ▓       |   Down:            37%
//     ▓       ░   Common Rank:     J-
//     |       ░   Efficiency Rank: B+
//       | | | ░   Source:          feedroll.com
//       ▓ ░ ▓ |
//       | | |
func BreakAway(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingBodyLong.avgPeriod + 4

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - settingBodyLong.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-4)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: long black (white)
	 * - second candle: black (white) day whose body gaps down (up)
	 * - third candle: black or white day with lower (higher) high and lower (higher) low than prior candle's
	 * - fourth candle: black (white) day with lower (higher) high and lower (higher) low than prior candle's
	 * - fifth candle: white (black) day that closes inside the gap, erasing the prior 3 days
	 * The meaning of "long" is specified with TA_SetCandleSettings
	 * outInteger is positive (1 to 100) when bullish or negative (-1 to -100) when bearish;
	 * the user should consider that breakaway is significant in a trend opposite to the last candle, while this
	 * function does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.realBody(i-4) > es.average(settingBodyLong, bodyLongPeriodTotal, i-4) && // 1st long
			es.candleColor(i-4) == es.candleColor(i-3) && // 1st, 2nd, 4th same color, 5th opposite
			es.candleColor(i-3) == es.candleColor(i-1) &&
			es.candleColor(i-1) == -es.candleColor(i) &&
			((es.candleColor(i-4) == -1 && // when 1st is black:
				es.realBodyGapDown(i-3, i-4) && // 2nd gaps down
				es.High(i-2) < es.High(i-3) && es.Low(i-2) < es.Low(i-3) && // 3rd has lower high and low than 2nd
				es.High(i-1) < es.High(i-2) && es.Low(i-1) < es.Low(i-2) && // 4th has lower high and low than 3rd
				es.Close(i) > es.Open(i-3) && es.Close(i) < es.Close(i-4)) || // 5th closes inside the gap

				(es.candleColor(i-4) == 1 && // when 1st is white:
					es.realBodyGapUp(i-3, i-4) && // 2nd gaps up
					es.High(i-2) > es.High(i-3) && es.Low(i-2) > es.Low(i-3) && // 3rd has higher high and low than 2nd
					es.High(i-1) > es.High(i-2) && es.Low(i-1) > es.Low(i-2) && // 4th has higher high and low than 3rd
					es.Close(i) < es.Open(i-3) && es.Close(i) > es.Close(i-4))) { // 5th closes inside the gap
			outInteger[i] = int(es.candleColor(i)) * 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-4) -
			es.rangeOf(settingBodyLong, bodyLongTrailingIdx-4)
		bodyLongTrailingIdx++
	}

	return outInteger
}
