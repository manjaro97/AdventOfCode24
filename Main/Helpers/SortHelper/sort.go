package SortHelper

import "sort"

func SortList(testList []int) []int {
	sort.Slice(testList, func(i, j int) bool {
		return testList[i] < testList[j]
	})

	return testList
}
