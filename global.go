package talibcdl

type candleSetting struct {
	rangeType rangeType
	avgPeriod int
	factor    float64
}

var (
	// real body is long when it's longer than the average of the 10 previous
	// candles' real body
	settingBodyLong = candleSetting{rangeTypeRealBody, 10, 1.0}
	// real body is very long when it's longer than 3 times the average of the 10
	// previous candles' real body
	settingBodyVeryLong = candleSetting{rangeTypeRealBody, 10, 3.0}
	// real body is short when it's shorter than the average of the 10 previous
	// candles' real bodies
	settingBodyShort = candleSetting{rangeTypeRealBody, 10, 1.0}
	// real body is like doji's body when it's shorter than 10% the average of the
	// 10 previous candles' high-low range
	settingBodyDoji = candleSetting{rangeTypeHighLow, 10, 0.1}
	// shadow is long when it's longer than the real body
	settingShadowLong = candleSetting{rangeTypeRealBody, 0, 1.0}
	// shadow is very long when it's longer than 2 times the real body
	settingShadowVeryLong = candleSetting{rangeTypeRealBody, 0, 2.0}
	// shadow is short when it's shorter than half the average of the 10 previous
	// candles' sum of shadows
	settingShadowShort = candleSetting{rangeTypeShadows, 10, 1.0}
	// shadow is very short when it's shorter than 10% the average of the 10
	// previous candles' high-low range
	settingShadowVeryShort = candleSetting{rangeTypeHighLow, 10, 0.1}
	// when measuring distance between parts of candles or width of gaps
	// "near" means "<= 20% of the average of the 5 previous candles' high-low range"
	settingNear = candleSetting{rangeTypeHighLow, 5, 0.2}
	// when measuring distance between parts of candles or width of gaps
	// "far" means ">= 60% of the average of the 5 previous candles' high-low range"
	settingFar = candleSetting{rangeTypeHighLow, 5, 0.6}
	// when measuring distance between parts of candles or width of gaps
	// "equal" means "<= 5% of the average of the 5 previous candles' high-low range"
	settingEqual = candleSetting{rangeTypeHighLow, 5, 0.05}
)

type rangeType int

const (
	rangeTypeRealBody rangeType = iota
	rangeTypeHighLow
	rangeTypeShadows
)

func (rt rangeType) rangeOf(s enhancedSeries, i int) float64 {
	switch rt {
	case rangeTypeRealBody:
		return s.realBody(i)
	case rangeTypeHighLow:
		return s.highLowRange(i)
	case rangeTypeShadows:
		return s.upperShadow(i) + s.lowerShadow(i)
	default:
		return 0
	}
}
