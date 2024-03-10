package s0006_zigzag_conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
}

func TestConvert2(t *testing.T) {
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
}
