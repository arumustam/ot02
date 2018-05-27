package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/arumustam/ot02/mycsv"
)

const outputDir = "output"

func main() {
	fmt.Println("reading")
	read_file, _ := os.OpenFile("./AddrBook.csv", os.O_RDONLY, 0644)
	_, dat := mycsv.Read(read_file)

	fmt.Println("sorting")
	start_t := time.Now()
	sort.Slice(dat, func(i, j int) bool {
		return dat[i][mycsv.LastNameKatakana] < dat[j][mycsv.LastNameKatakana]
	})

	sorting_time := time.Now().Sub(start_t).Seconds()
	fmt.Println("***sorting time:", sorting_time*(1000), "ms")
}
