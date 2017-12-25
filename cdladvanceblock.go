package talibcdl

import (
	"log"
)

// AdvanceBlock implements ta-lib function TA_CDLADVANCEBLOCK.
//
//       | |   Up:              64%
//       | ░   Down:            36%
//     | ░ |   Common Rank:     G
//     ░ |     Efficiency Rank: F
//     ░       Source:          feedroll.com
//     ░
func AdvanceBlock(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := intMax(
		intMax(
			intMax(settingShadowLong.avgPeriod, settingShadowShort.avgPeriod),
			intMax(settingFar.avgPeriod, settingNear.avgPeriod)),
		settingBodyLong.avgPeriod) + 2

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.
	shadowShortTrailingIdx := startIdx - settingShadowShort.avgPeriod
	shadowLongTrailingIdx := startIdx - settingShadowLong.avgPeriod
	nearTrailingIdx := startIdx - settingNear.avgPeriod
	farTrailingIdx := startIdx - settingFar.avgPeriod
	bodyLongPeriodTotal := 0.0
	bodyLongTrailingIdx := startIdx - settingBodyLong.avgPeriod

	shadowShortPeriodTotal := [3]float64{}
	shadowLongPeriodTotal := [2]float64{}
	nearPeriodTotal := [3]float64{}
	farPeriodTotal := [3]float64{}

	for i := shadowShortTrailingIdx; i < startIdx; i++ {
		shadowShortPeriodTotal[2] += es.rangeOf(settingShadowShort, i-2)
		shadowShortPeriodTotal[1] += es.rangeOf(settingShadowShort, i-1)
		shadowShortPeriodTotal[0] += es.rangeOf(settingShadowShort, i)
	}
	for i := shadowLongTrailingIdx; i < startIdx; i++ {
		shadowLongPeriodTotal[1] += es.rangeOf(settingShadowLong, i-1)
		shadowLongPeriodTotal[0] += es.rangeOf(settingShadowLong, i)
	}
	for i := nearTrailingIdx; i < startIdx; i++ {
		nearPeriodTotal[2] += es.rangeOf(settingNear, i-2)
		nearPeriodTotal[1] += es.rangeOf(settingNear, i-1)
	}
	for i := farTrailingIdx; i < startIdx; i++ {
		farPeriodTotal[2] += es.rangeOf(settingFar, i-2)
		farPeriodTotal[1] += es.rangeOf(settingFar, i-1)
	}
	for i := bodyLongTrailingIdx; i < startIdx; i++ {
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-2)
	}

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - three white candlesticks with consecutively higher closes
	 * - each candle opens within or near the previous white real body
	 * - first candle: long white with no or very short upper shadow (a short shadow is accepted too for more flexibility)
	 * - second and third candles, or only third candle, show signs of weakening: progressively smaller white real bodies
	 * and/or relatively long upper shadows; see below for specific conditions
	 * The meanings of "long body", "short shadow", "far" and "near" are specified with TA_SetCandleSettings;
	 * outInteger is negative (-1 to -100): advance block is always bearish;
	 * the user should consider that advance block is significant when it appears in uptrend, while this function
	 * does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if es.candleColor(i-2) == 1 && // 1st white
			es.candleColor(i-1) == 1 && // 2nd white
			es.candleColor(i) == 1 && // 3rd white
			es.Close(i) > es.Close(i-1) && es.Close(i-1) > es.Close(i-2) && // consecutive higher closes
			es.Open(i-1) > es.Open(i-2) && // 2nd opens within/near 1st real body
			es.Open(i-1) <= es.Close(i-2)+es.average(settingNear, nearPeriodTotal[2], i-2) &&
			es.Open(i) > es.Open(i-1) && // 3rd opens within/near 2nd real body
			es.Open(i) <= es.Close(i-1)+es.average(settingNear, nearPeriodTotal[1], i-1) &&
			es.realBody(i-2) > es.average(settingBodyLong, bodyLongPeriodTotal, i-2) && // 1st: long real body
			es.upperShadow(i-2) < es.average(settingShadowShort, shadowShortPeriodTotal[2], i-2) &&
			// 1st: short upper shadow
			(
			// ( 2 far smaller than 1 && 3 not longer than 2 )
			// advance blocked with the 2nd, 3rd must not carry on the advance
			(es.realBody(i-1) < es.realBody(i-2)-es.average(settingFar, farPeriodTotal[2], i-2) &&
				es.realBody(i) < es.realBody(i-1)+es.average(settingNear, nearPeriodTotal[1], i-1)) ||
				// 3 far smaller than 2
				// advance blocked with the 3rd
				(es.realBody(i) < es.realBody(i-1)-es.average(settingFar, farPeriodTotal[1], i-1)) ||
				// ( 3 smaller than 2 && 2 smaller than 1 && (3 or 2 not short upper shadow) )
				// advance blocked with progressively smaller real bodies and some upper shadows
				(es.realBody(i) < es.realBody(i-1) &&
					es.realBody(i-1) < es.realBody(i-2) &&
					(es.upperShadow(i) > es.average(settingShadowShort, shadowShortPeriodTotal[0], i) ||
						es.upperShadow(i-1) > es.average(settingShadowShort, shadowShortPeriodTotal[1], i-1))) ||
				// ( 3 smaller than 2 && 3 long upper shadow )
				// advance blocked with 3rd candle's long upper shadow and smaller body
				(es.realBody(i) < es.realBody(i-1) &&
					es.upperShadow(i) > es.average(settingShadowLong, shadowLongPeriodTotal[0], i))) {
			outInteger[i] = -100
		}

		// add the current range and subtract the first range: this is done after the pattern recognition
		// when avgPeriod is not 0, that means "compare with the previous candles" (it excludes the current candle)
		for totIdx := 2; totIdx >= 0; totIdx-- {
			shadowShortPeriodTotal[totIdx] += es.rangeOf(settingShadowShort, i-totIdx) -
				es.rangeOf(settingShadowShort, shadowShortTrailingIdx-totIdx)
		}
		for totIdx := 1; totIdx >= 0; totIdx-- {
			shadowLongPeriodTotal[totIdx] += es.rangeOf(settingShadowLong, i-totIdx) -
				es.rangeOf(settingShadowLong, shadowLongTrailingIdx-totIdx)
		}
		for totIdx := 2; totIdx >= 1; totIdx-- {
			farPeriodTotal[totIdx] += es.rangeOf(settingFar, i-totIdx) -
				es.rangeOf(settingFar, farTrailingIdx-totIdx)
			nearPeriodTotal[totIdx] += es.rangeOf(settingNear, i-totIdx) -
				es.rangeOf(settingNear, nearTrailingIdx-totIdx)
		}
		bodyLongPeriodTotal += es.rangeOf(settingBodyLong, i-2) - es.rangeOf(settingBodyLong, bodyLongTrailingIdx-2)
		shadowShortTrailingIdx++
		shadowLongTrailingIdx++
		nearTrailingIdx++
		farTrailingIdx++
		bodyLongTrailingIdx++
	}

	return outInteger
}
