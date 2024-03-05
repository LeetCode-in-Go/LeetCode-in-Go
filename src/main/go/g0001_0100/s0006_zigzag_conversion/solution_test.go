package s0006_zigzag_conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
}

func TestConvert2(t *testing.T) {
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
}
