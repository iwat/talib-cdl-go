package talibcdl

import "log"

// Doji implements ta-lib function TA_CDLDOJI.
//
//     +       Up:              43%
//             Down:            57%
//             Common Rank:     E-
//             Efficiency Rank: J+
//             Source:          feedroll.com
func Doji(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := settingBodyDoji.avgPeriod

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyDojiPeriodTotal := 0.0
	bodyDojiTrailingIdx := startIdx - settingBodyDoji.avgPeriod

	for i := bodyDojiTrailingIdx; i < startIdx; i++ {
		bodyDojiPeriodTotal += es.rangeOf(settingBodyDoji, i)
	}

	/* Proceed with the calculation for the requested range.
	 *
	 * Must have:
	 * - open quite equal to close
	 * How much can be the maximum distance between open and close is specified with TA_SetCandleSettings
	 * outInteger is always positive (1 to 100) but this does not mean it is bullish: doji shows uncertainty and it is
	 * neither bullish nor bearish when considered alone
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.realBody(i) <= es.average(settingBodyDoji, bodyDojiPeriodTotal, i) {
			outInteger[i] = 100
		} else {
			outInteger[i] = 0
		}
		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyDojiPeriodTotal += es.rangeOf(settingBodyDoji, i) - es.rangeOf(settingBodyDoji, bodyDojiTrailingIdx)
		bodyDojiTrailingIdx++
	}

	return outInteger
}
