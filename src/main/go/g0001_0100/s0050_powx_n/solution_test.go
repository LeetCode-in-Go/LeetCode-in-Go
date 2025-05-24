package s0050_powx_n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyPow(t *testing.T) {
	assert.InDelta(t, 1024.00000, myPow(2.00000, 10), 1e-9)
}

func TestMyPow2(t *testing.T) {
	assert.InDelta(t, 9.261000000000001, myPow(2.10000, 3), 1e-9)
}

func TestMyPow3(t *testing.T) {
	assert.InDelta(t, 0.25000, myPow(2.00000, -2), 1e-9)
}
