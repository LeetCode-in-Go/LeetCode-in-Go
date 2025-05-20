package s0135_candy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCandy(t *testing.T) {
	ratings := []int{1, 0, 2}
	assert.Equal(t, 5, candy(ratings))
}

func TestCandy2(t *testing.T) {
	ratings := []int{1, 2, 2}
	assert.Equal(t, 4, candy(ratings))
}

func TestCandy3(t *testing.T) {
	ratings := []int{1, 3, 2, 2, 1}
	assert.Equal(t, 7, candy(ratings))
}

func TestCandy4(t *testing.T) {
	ratings := []int{1, 2, 87, 87, 87, 2, 1}
	assert.Equal(t, 13, candy(ratings))
}

func TestCandy5(t *testing.T) {
	ratings := []int{1}
	assert.Equal(t, 1, candy(ratings))
}

func TestCandy6(t *testing.T) {
	var ratings []int
	assert.Equal(t, 0, candy(ratings))
}
