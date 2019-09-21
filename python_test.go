package talibcdl

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func compareInts(t *testing.T, series SimpleSeries, goResult []int, taCall string) {
	t.Helper()

	pyprog := fmt.Sprintf(`import talib,numpy
testOpen = numpy.array(%s)
testHigh = numpy.array(%s)
testLow = numpy.array(%s)
testClose = numpy.array(%s)
testVolume = numpy.array(%s)
testRand = numpy.array(%s)
%s
print(' '.join([str(p) for p in result]).replace('nan','0.0'))`,
		a2s(series.Opens), a2s(series.Highs), a2s(series.Lows), a2s(series.Closes), a2s(series.Volumes), a2s(series.Rands), taCall)

	pyOut, err := exec.Command("python3", "-c", pyprog).CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %v\n%s", err, string(pyOut))
	}

	var pyResult []int
	strResult := strings.Fields(string(pyOut))
	for _, arg := range strResult {
		if n, err := strconv.Atoi(arg); err == nil {
			pyResult = append(pyResult, n)
		}
	}

	if len(goResult) != len(pyResult) {
		t.Fatalf("different size\ngo: %#v\npy: %#v\n", len(goResult), len(pyResult))
	}

	for i := 0; i < len(goResult); i++ {
		if goResult[i] != pyResult[i] {
			t.Fatalf("index %d mismatch\ngo: %#v\npy: %#v\ngo full: %v\npy full: %v", i, goResult[i], pyResult[i], goResult, pyResult)
		}
	}
}

func a2s(a []float64) string { // go float64 array to python list initializer string
	return strings.Replace(fmt.Sprintf("%f", a), " ", ",", -1)
}
