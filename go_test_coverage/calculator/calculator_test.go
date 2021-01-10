package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type in struct {
	a uint8
	b uint8
}

type Test struct {
	name string
	in   in
	out  uint16
}

var tests = []Test{
	{"small numbers", in{1, 2}, 3},
	{"edge at the range", in{127, 127}, 254},
	{"over the range", in{130, 130}, 260},
}

func TestSum(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.out, sum(test.in.a, test.in.b), fmt.Sprintf("%s: %v", test.name, test.in))
	}
}
