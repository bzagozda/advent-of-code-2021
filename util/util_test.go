package util

import (
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	data := []int{1, 2, 3}
	got := SlidingWindow(data, 2)
	Assert(t, len(got), 2)
	Assert(t, got[0], []int{1, 2})
	Assert(t, got[1], []int{2, 3})

	got = SlidingWindow(data, 0)
	Assert(t, len(got), 0)

	got = SlidingWindow([]int{}, 2)
	Assert(t, len(got), 0)
}

func TestCount(t *testing.T) {
	data := []bool{true, false, true}
	got := Count(3, func(i int) bool { return data[i] })
	Assert(t, got, 2)

	got = Count(0, func(i int) bool { return data[i] })
	Assert(t, got, 0)
}

func TestMapToInt(t *testing.T) {
	data := []string{}
	got := MapToInt(data)
	Assert(t, got, []int{})

	data = []string{"1", "2", "3"}
	got = MapToInt(data)
	Assert(t, got, []int{1, 2, 3})

	data = []string{"a", "2", "3"}
	got = MapToInt(data)
	Assert(t, got, []int{2, 3})
}

func TestArraySum(t *testing.T) {
	data := []int{1, 2, 3}
	got := ArraySum(data)
	Assert(t, got, 6)

	data = []int{}
	got = ArraySum(data)
	Assert(t, got, 0)

	data = []int{-1, 1}
	got = ArraySum(data)
	Assert(t, got, 0)
}
