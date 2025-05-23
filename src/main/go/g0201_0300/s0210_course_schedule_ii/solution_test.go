package s0210_course_schedule_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCourseScheduleII(t *testing.T) {
	prerequisites := [][]int{{1, 0}}
	numCourses := 2
	assert.Equal(t, []int{0, 1}, findOrder(numCourses, prerequisites))
}

func TestCourseScheduleII2(t *testing.T) {
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	numCourses := 4
	assert.Equal(t, []int{0, 1, 2, 3}, findOrder(numCourses, prerequisites))
}

func TestCourseScheduleII3(t *testing.T) {
	prerequisites := [][]int{}
	numCourses := 1
	assert.Equal(t, []int{0}, findOrder(numCourses, prerequisites))
}
