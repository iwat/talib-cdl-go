package talibcdl

import (
	"log"
)

// ThreeWhiteSoldiers implements ta-lib function TA_CDL3WHITESOLDIERS.
//
//         ░
//       ░ ░
//     ░ ░
//     ░
func ThreeWhiteSoldiers(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := intMax(
		intMax(settingShadowVeryShort.avgPeriod, settingBodyShort.avgPeriod),
		intMax(settingFar.avgPeriod, settingNear.avgPeriod)) + 2

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	shadowVeryShortTrailingIdx := startIdx - settingShadowVeryShort.avgPeriod
	nearTrailingIdx := startIdx - settingNear.avgPeriod
	farTrailingIdx := startIdx - settingFar.avgPeriod
	bodyShortPeriodTotal := 0.0
	bodyShortTrailingIdx := startIdx - settingBodyShort.avgPeriod

	shadowVeryShortPeriodTotal := [3]float64{}
	nearPeriodTotal := [3]float64{}
	farPeriodTotal := [3]float64{}
	for i := shadowVeryShortTrailingIdx; i < startIdx; i++ {
		shadowVeryShortPeriodTotal[2] += es.rangeOf(settingShadowVeryShort, i-2)
		shadowVeryShortPeriodTotal[1] += es.rangeOf(settingShadowVeryShort, i-1)
		shadowVeryShortPeriodTotal[0] += es.rangeOf(settingShadowVeryShort, i)
	}
	for i := nearTrailingIdx; i < startIdx; i++ {
		nearPeriodTotal[2] += es.rangeOf(settingNear, i-2)
		nearPeriodTotal[1] += es.rangeOf(settingNear, i-1)
	}
	for i := farTrailingIdx; i < startIdx; i++ {
		farPeriodTotal[2] += es.rangeOf(settingFar, i-2)
		farPeriodTotal[1] += es.rangeOf(settingFar, i-1)
	}
	for i := bodyShortTrailingIdx; i < startIdx; i++ {
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - three white candlesticks with consecutively higher closes
	 * - Greg Morris wants them to be long, Steve Nison doesn't; anyway they should not be short
	 * - each candle opens within or near the previous white real body
	 * - each candle must have no or very short upper shadow
	 * - to differentiate this pattern from advance block, each candle must not be far shorter than the prior candle
	 * The meanings of "not short", "very short shadow", "far" and "near" are specified with TA_SetCandleSettings;
	 * here the 3 candles must be not short, if you want them to be long use TA_SetCandleSettings on BodyShort;
	 * outInteger is positive (1 to 100): advancing 3 white soldiers is always bullish;
	 * the user should consider that 3 white soldiers is significant when it appears in downtrend, while this function
	 * does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-2).isWhite() && // 1st white
			es.upperShadow(i-2) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[2], i-2) &&
			// very short upper shadow
			es.candleColor(i-1).isWhite() && // 2nd white
			es.upperShadow(i-1) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[1], i-1) &&
			// very short upper shadow
			es.candleColor(i).isWhite() && // 3rd white
			es.upperShadow(i) < es.average(settingShadowVeryShort, shadowVeryShortPeriodTotal[0], i) &&
			// very short upper shadow
			es.Close(i) > es.Close(i-1) && es.Close(i-1) > es.Close(i-2) && // consecutive higher closes
			es.Open(i-1) > es.Open(i-2) && // 2nd opens within/near 1st real body
			es.Open(i-1) <= es.Close(i-2)+es.average(settingNear, nearPeriodTotal[2], i-2) &&
			es.Open(i) > es.Open(i-1) && // 3rd opens within/near 2nd real body
			es.Open(i) <= es.Close(i-1)+es.average(settingNear, nearPeriodTotal[1], i-1) &&
			es.realBody(i-1) > es.realBody(i-2)-es.average(settingFar, farPeriodTotal[2], i-2) &&
			// 2nd not far shorter than 1st
			es.realBody(i) > es.realBody(i-1)-es.average(settingFar, farPeriodTotal[1], i-1) &&
			// 3rd not far shorter than 2nd
			es.realBody(i) > es.average(settingBodyShort, bodyShortPeriodTotal, i) { // not short real body
			outInteger[i] = 100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		for totIdx := 2; totIdx >= 0; totIdx-- {
			shadowVeryShortPeriodTotal[totIdx] += es.rangeOf(settingShadowVeryShort, i-totIdx) - es.rangeOf(settingShadowVeryShort, shadowVeryShortTrailingIdx-totIdx)
		}
		for totIdx := 2; totIdx >= 1; totIdx-- {
			farPeriodTotal[totIdx] += es.rangeOf(settingFar, i-totIdx) - es.rangeOf(settingFar, farTrailingIdx-totIdx)
			nearPeriodTotal[totIdx] += es.rangeOf(settingNear, i-totIdx) - es.rangeOf(settingNear, nearTrailingIdx-totIdx)
		}
		bodyShortPeriodTotal += es.rangeOf(settingBodyShort, i) - es.rangeOf(settingBodyShort, bodyShortTrailingIdx)
		shadowVeryShortTrailingIdx++
		nearTrailingIdx++
		farTrailingIdx++
		bodyShortTrailingIdx++
	}

	return outInteger
}
