package sort

func MergeSort(cmp func(a []string, b []string) bool, dat [][]string) [][]string {
	if len(dat) <= 1 {
		return dat
	}

	mid := len(dat) / 2
	left_dat := dat[:mid]
	right_dat := dat[mid:]

	left_dat = MergeSort(cmp, left_dat)
	right_dat = MergeSort(cmp, right_dat)

	return merge(cmp, left_dat, right_dat)
}

func GoMergeSort(cmp func(a []string, b []string) bool, dat [][]string) [][]string {
	if len(dat) <= 1 {
		return dat
	}

	mid := len(dat) / 2
	left_dat := dat[:mid]
	right_dat := dat[mid:]

	if len(left_dat) > 8192 /* 2の乗数で実験した結果、この数字が安定している */ {
		lch := make(chan bool)
		rch := make(chan bool)
		go func() {
			left_dat = GoMergeSort(cmp, left_dat)
			lch <- true
		}()
		go func() {
			right_dat = GoMergeSort(cmp, right_dat)
			rch <- true
		}()
		<-lch
		<-rch
	} else {
		left_dat = MergeSort(cmp, left_dat)
		right_dat = MergeSort(cmp, right_dat)
	}
	return merge(cmp, left_dat, right_dat)
}

func merge(cmp func(a []string, b []string) bool, left_dat [][]string, right_dat [][]string) [][]string {
	var ret [][]string
	li := 0
	ri := 0

	for li < len(left_dat) && ri < len(right_dat) {
		if cmp(left_dat[li], right_dat[ri]) {
			ret = append(ret, left_dat[li])
			li++
		} else {
			ret = append(ret, right_dat[ri])
			ri++
		}
	}
	if len(left_dat) != 0 {
		ret = append(ret, left_dat[li:]...)
	}
	if len(right_dat) != 0 {
		ret = append(ret, right_dat[ri:]...)
	}
	return ret
}
