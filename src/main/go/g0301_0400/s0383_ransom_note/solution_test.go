package s0383_ransom_note

import (
	"github.com/stretchr/testify/assert"
    "testing"
)

func TestCanConstruct(t *testing.T) {
	assert.False(t, canConstruct("a", "b"))
}

func TestCanConstruct2(t *testing.T) {
	assert.False(t, canConstruct("aa", "ab"))
}

func TestCanConstruct3(t *testing.T) {
	assert.True(t, canConstruct("aa", "aab"))
}

func TestCanConstruct4(t *testing.T) {
	assert.True(t, canConstruct("", ""))
}
