package s0205_isomorphic_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	assert.True(t, isIsomorphic("egg", "add"))
}

func TestIsIsomorphic2(t *testing.T) {
	assert.False(t, isIsomorphic("foo", "bar"))
}

func TestIsIsomorphic3(t *testing.T) {
	assert.True(t, isIsomorphic("paper", "title"))
}
