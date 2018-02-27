package special

import "math"

const machineEpsilon = 1. / (2 << 52) // 2**-53

const maxLog = 1024 * math.Ln2 // log(2**1024)

const minLog = -1022 * math.Ln2 // log(2**-1022)

const testDelta = 1. / (2 << 40)
