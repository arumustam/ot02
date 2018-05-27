package sort

func QuickSort(cmp func(a []string, b []string) bool, dat [][]string) [][]string {
	if len(dat) < 1 {
		return dat
	}
	pivot := dat[0]
	var left [][]string
	var right [][]string

	for i := 1; i < len(dat); i++ {
		if cmp(dat[i], pivot) {
			left = append(left, dat[i])
		} else {
			right = append(right, dat[i])
		}
	}
	left = QuickSort(cmp, left)
	right = QuickSort(cmp, right)
	left = append(left, pivot)
	return append(left, right...)
}
