package talibcdl

import (
	"log"
)

// ThreeOutside implements ta-lib function TA_CDL3OUTSIDE.
//
//       |
//     | ▓ |
//     ░ ▓ ▓
//     ░ ▓ ▓
//     | ▓ ▓
//       | ▓
//         |
func ThreeOutside(series Series) []int {
	es := enhancedSeries{series}
	outInteger := make([]int, es.Len())

	// Identify the minimum number of price bar needed to calculate at least one output.
	// Move up the start index if there is not enough initial data.
	startIdx := 3

	// Make sure there is still something to evaluate.
	if startIdx >= es.Len() {
		log.Printf("too few input len(%d) want(%d)", es.Len(), startIdx)
		return outInteger
	}

	// Do the calculation using tight loops.
	// Add-up the initial period, except for the last value.

	/* Proceed with the calculation for the requested range.
	 * Must have:
	 * - first: black (white) real body
	 * - second: white (black) real body that engulfs the prior real body
	 * - third: candle that closes higher (lower) than the second candle
	 * outInteger is positive (1 to 100) for the three outside up or negative (-1 to -100) for the three outside down;
	 * the user should consider that a three outside up must appear in a downtrend and three outside down must appear
	 * in an uptrend, while this function does not consider it
	 */
	for i := startIdx; i < es.Len(); i++ {
		if (es.candleColor(i-1) == 1 && es.candleColor(i-2) == -1 && // white engulfs black
			es.Close(i-1) > es.Open(i-2) && es.Open(i-1) < es.Close(i-2) &&
			es.Close(i) > es.Close(i-1)) || // third candle higher
			(es.candleColor(i-1) == -1 && es.candleColor(i-2) == 1 && // black engulfs white
				es.Open(i-1) > es.Close(i-2) && es.Close(i-1) < es.Open(i-2) &&
				es.Close(i) < es.Close(i-1)) { // third candle lower
			outInteger[i] = int(es.candleColor(i-1)) * 100
		}
	}

	return outInteger
}
