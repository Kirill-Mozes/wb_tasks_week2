package funcs

import "sort"

func simpleSort(arr []string, rev bool) []string {
	if rev {
		sort.Sort(sort.Reverse(sort.StringSlice(arr)))
	} else {
		sort.Strings(arr)
	}
	return arr

}
