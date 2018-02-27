package special

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogistic(t *testing.T) {
	for x := -6.; x <= 6; x += 0.1 {
		t.Run(fmt.Sprintf("x=%.1f", x), func(t *testing.T) {
			assert.InDelta(t, 1., Logistic(x)+Logistic(-x), testDelta)
		})
	}
}
