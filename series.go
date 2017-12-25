package talibcdl

import (
	"math"
)

type Series interface {
	Len() int
	High(i int) float64
	Open(i int) float64
	Close(i int) float64
	Low(i int) float64
}

type SimpleSeries struct {
	Highs   []float64
	Opens   []float64
	Closes  []float64
	Lows    []float64
	Volumes []float64
	Rands   []float64
}

func (s SimpleSeries) Len() int {
	return len(s.Highs)
}

func (s SimpleSeries) High(i int) float64 {
	return s.Highs[i]
}

func (s SimpleSeries) Open(i int) float64 {
	return s.Opens[i]
}

func (s SimpleSeries) Close(i int) float64 {
	return s.Closes[i]
}

func (s SimpleSeries) Low(i int) float64 {
	return s.Lows[i]
}

type enhancedSeries struct {
	Series
}

func (s enhancedSeries) average(st candleSetting, sum float64, i int) float64 {
	a := s.rangeOf(st, i)
	if st.avgPeriod != 0.0 {
		a = sum / float64(st.avgPeriod)
	}
	b := 1.0
	if st.rangeType == rangeTypeShadows {
		b = 2.0
	}
	return st.factor * a / b
}

func (s enhancedSeries) candleColor(i int) candleColor {
	if s.Close(i) >= s.Open(i) {
		return candleColorWhite
	} else {
		return candleColorBlack
	}
}

func (s enhancedSeries) highLowRange(i int) float64 {
	return s.High(i) - s.Low(i)
}

func (s enhancedSeries) isCandleGapDown(i1, i2 int) bool {
	return s.High(i1) < s.Low(i2)
}

func (s enhancedSeries) isCandleGapUp(i1, i2 int) bool {
	return s.Low(i1) > s.High(i2)
}

func (s enhancedSeries) lowerShadow(i int) float64 {
	return math.Min(s.Close(i), s.Open(i)) - s.Low(i)
}

func (s enhancedSeries) rangeOf(st candleSetting, i int) float64 {
	return st.rangeType.rangeOf(s, i)
}

func (s enhancedSeries) realBody(i int) float64 {
	return math.Abs(s.Close(i) - s.Open(i))
}

func (s enhancedSeries) realBodyGapDown(i2, i1 int) bool {
	return math.Max(s.Open(i2), s.Close(i2)) < math.Min(s.Open(i1), s.Close(i1))
}

func (s enhancedSeries) realBodyGapUp(i2, i1 int) bool {
	return math.Min(s.Open(i2), s.Close(i2)) > math.Max(s.Open(i1), s.Close(i1))
}

func (s enhancedSeries) upperShadow(i int) float64 {
	return s.High(i) - (math.Max(s.Close(i), s.Open(i)))
}

type candleColor int

const (
	candleColorWhite candleColor = 1
	candleColorBlack             = -1
)

func (c candleColor) isBlack() bool {
	return c == candleColorBlack
}

func (c candleColor) isWhite() bool {
	return c == candleColorWhite
}
