package special

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncompleteGamma(t *testing.T) {
	for a := testDelta; a <= 6; a += 0.1 {
		for x := testDelta; x <= 6; x += 0.1 {
			t.Run(fmt.Sprintf("a=%.1f,x=%.1f", a, x), func(t *testing.T) {
				assert.InDelta(t, 1., LowerIncompleteGamma(a, x)+UpperIncompleteGamma(a, x), testDelta)
			})
		}
	}
}
