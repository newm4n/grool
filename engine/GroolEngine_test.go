package engine

import (
	"sort"
	"testing"
)

type Sorting struct {
	Val int
}

func TestGroolSorting(t *testing.T) {
	arr := make([]*Sorting, 0)
	arr = append(arr, &Sorting{Val: 4})
	arr = append(arr, &Sorting{Val: 7})
	arr = append(arr, &Sorting{Val: 3})
	arr = append(arr, &Sorting{Val: 6})
	arr = append(arr, &Sorting{Val: 9})
	arr = append(arr, &Sorting{Val: 8})
	arr = append(arr, &Sorting{Val: 1})
	arr = append(arr, &Sorting{Val: 2})

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val > arr[j].Val
	})

	if arr[0].Val != 9 {
		t.FailNow()
	}
}

func TestGrool_Execute(t *testing.T) {

}
