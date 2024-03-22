package s0207_course_schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanFinish(t *testing.T) {
	assert.Equal(t, true, canFinish(2, [][]int{{1, 0}}))
}

func TestCanFinish2(t *testing.T) {
	assert.Equal(t, false, canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
