package s0069_sqrtx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 2, mySqrt(4))
}

func TestMySqrt2(t *testing.T) {
	assert.Equal(t, 2, mySqrt(8))
}
