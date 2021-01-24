package surface

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareSurface(t *testing.T) {
	r := square(5)
	assert.Equal(t, 25, r)
}
