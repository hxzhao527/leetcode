package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	data := []struct {
		N1, N2 []int
		Medium float64
	}{
		{
			N1:     []int{1, 3},
			N2:     []int{2},
			Medium: 2.0,
		},
		{
			N1:     []int{1, 2},
			N2:     []int{3, 4},
			Medium: 2.5,
		},
	}
	for _, d := range data {
		if ret := FindMedianSortedArrays(d.N1, d.N2); ret != d.Medium {
			t.Errorf("FindMedianSortedArrays %v, %v, expect %f, got %f", d.N1, d.N2, d.Medium, ret)
		}
	}
}
