package sort

func MakeCmp(nCols []int, orders []bool) func(a []string, b []string) bool {
	return func(a []string, b []string) bool {
		for i := 0; i < len(nCols); i++ {
			if a[nCols[i]] < b[nCols[i]] {
				return orders[i]
			} else if a[nCols[i]] > b[nCols[i]] {
				return !orders[i]
			}
		}
		return true
	}
}
