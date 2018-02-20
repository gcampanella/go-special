package special

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const epsilon = 1e-9

func TestLogistic(t *testing.T) {
	for x := -6.; x <= 6; x += 0.1 {
		assert.InEpsilon(t, 1., Logistic(x)+Logistic(-x), epsilon)
	}
}
