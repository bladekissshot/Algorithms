package Sorting

import "testing"

func TestQuickSort(t *testing.T) {
	if ans := QuickSort(); ans != "abc" {
		t.Errorf("error")
	}
}
