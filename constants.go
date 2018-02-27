package special

import "math"

const (
	machineEpsilon = 1. / (2 << 52) // 2**-53

	maxLog = 1024 * math.Ln2 // log(2**1024)

	minLog = -1022 * math.Ln2 // log(2**-1022)

	testDelta = 1. / (2 << 40)
)
