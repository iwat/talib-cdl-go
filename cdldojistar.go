package talibcdl

import "log"

// DojiStart implements ta-lib function TA_CDLDOJISTAR.
//
//     |       Up:              36%
//     ▓       Down:            64%
//     ▓       Common Rank:     F+
//     |       Efficiency Rank: E-
//      +      Source:          feedroll.com
func DojiStar(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	startIdx := intMax(settingBodyDoji.avgPeriod, settingBodyLong.avgPeriod) + 1

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	bodyLongPeriodTotal := 0.0
	bodyDojiPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - 1 - settingBodyLong.avgPeriod
	bodyDojiTrailingIdx := startIdx - settingBodyDoji.avgPeriod

	for i := bodyLongTrailingIdx; i < startIdx-1; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i)
	}

	for i := bodyDojiTrailingIdx; i < startIdx; i++ {
		bodyDojiPeriodTotal += es.rangeOf(settingBodyDoji, i)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first candle: long real body
	 * - second candle: star (open gapping up in an uptrend or down in a downtrend) with a doji
	 * The meaning of "doji" and "long" is specified with TA_SetCandleSettings
	 * outInteger is positive (1 to 100) when bullish or negative (-1 to -100) when bearish;
	 * it's defined bullish when the long candle is white and the star gaps up, bearish when the long candle
	 * is black and the star gaps down; the user should consider that a doji star is bullish when it appears
	 * in an uptrend and it's bearish when it appears in a downtrend, so to determine the bullishness or
	 * bearishness of the pattern the trend must be analyzed
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.realBody(i-1) > es.average(settingBodyLong, bodyLongPeriodTotal, i-1) && // 1st: long real body
			es.realBody(i) <= es.average(settingBodyDoji, bodyDojiPeriodTotal, i) && // 2nd: doji
			((es.candleColor(i-1).isWhite() && es.realBodyGapUp(i, i-1)) || // that gaps up if 1st is white
				(es.candleColor(i-1).isBlack() && es.realBodyGapDown(i, i-1))) { // or down if 1st is black
			outInteger[i] = -int(es.candleColor(i-1)) * 100
		} else {
			outInteger[i] = 0
		}
		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-1) - es.rangeOf(settingBodyLong, bodyLongTrailingIdx)
		bodyDojiPeriodTotal += es.rangeOf(settingBodyDoji, i) - es.rangeOf(settingBodyDoji, bodyDojiTrailingIdx)
		bodyLongTrailingIdx++
		bodyDojiTrailingIdx++
	}

	return outInteger
}
