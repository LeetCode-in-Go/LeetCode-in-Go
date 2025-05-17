package s0071_simplify_path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	assert.Equal(t, "/home", simplifyPath("/home/"))
}

func TestSimplifyPath2(t *testing.T) {
	assert.Equal(t, "/", simplifyPath("/../"))
}

func TestSimplifyPath3(t *testing.T) {
	assert.Equal(t, "/home/foo", simplifyPath("/home//foo/"))
}
