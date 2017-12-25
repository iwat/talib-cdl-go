package talibcdl

import (
	"fmt"
	"math"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func compareInts(t *testing.T, series SimpleSeries, goResult []int, taCall string) {
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

	pyOut, err := exec.Command("python", "-c", pyprog).Output()
	ok(t, err)

	var pyResult []int
	strResult := strings.Fields(string(pyOut))
	for _, arg := range strResult {
		if n, err := strconv.Atoi(arg); err == nil {
			pyResult = append(pyResult, n)
		}
	}

	equals(t, len(goResult), len(pyResult))

	for i := 0; i < len(goResult); i++ {
		if goResult[i] != pyResult[i] {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("%s:%d:\n\tgo!: %#v\n\tpy!: %#v\n", filepath.Base(file), line, goResult[i], pyResult[i])
			t.FailNow()
		}
	}
}

func a2s(a []float64) string { // go float64 array to python list initializer string
	return strings.Replace(fmt.Sprintf("%f", a), " ", ",", -1)
}

func ok(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n", filepath.Base(file), line, err.Error())
		t.FailNow()
	}
}

func equals(t *testing.T, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\tgo: %#v\n\tpy: %#v\n", filepath.Base(file), line, exp, act)
		t.FailNow()
	}
}

func round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}
