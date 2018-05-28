package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"

	"github.com/arumustam/ot02/mycsv"
	"github.com/arumustam/ot02/sort"
	"github.com/arumustam/ot02/utils"
)

const outputDirName = "output_sortapp"

func makeSortApp() *cli.App {
	app := cli.NewApp()
	app.Name = "sortapp"
	app.Usage = "sortapp [AddressBookPath]"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "columns, c",
			Value: "1,2,3",
			Usage: "ソートしたいカラムの番号を指定。番号は左から1,2,3,...と割り当てられている。",
		},
		cli.StringFlag{
			Name:  "orders, o",
			Value: "a,d,a",
			Usage: "昇順(ascending)か降順(descending)を指定。columnsフラグで指定したカラムと位置関係が対応している。",
		},
	}
	return app
}

func main() {
	exe, _ := os.Executable()
	exepath := filepath.Dir(exe)
	outputDir := filepath.Join(exepath, outputDirName)

	// outputはカレントディクトリの"sortapp_output"ディレクトリで行われる
	if !utils.Exists(outputDir) {
		_ = os.Mkdir(outputDir, 0777)
	}

	app := makeSortApp()
	app.Action = func(c *cli.Context) error {
		var addrBookPath string
		if len(c.Args()) == 1 {
			addrBookPath = c.Args()[0]
		} else {
			fmt.Println("アドレス帳のパスを指定してください。")
			os.Exit(1)
		}
		addrBookName := filepath.Base(addrBookPath)
		sortedAddrBookPath := utils.GetNewFilePath(addrBookName, outputDir, 2)

		// ******************** reading ********************
		read_file, err := os.OpenFile(addrBookPath, os.O_RDONLY, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		head, dat := mycsv.Read(read_file)

		// ******************** sorting ********************
		columns := strings.Split(c.String("columns"), ",")
		orders := strings.Split(c.String("orders"), ",")

		if len(columns) != len(orders) {
			fmt.Println("columnsとordersの数が合っていません。")
			os.Exit(1)
		}

		// ********** columns
		nCols := []int{}
		for _, col := range columns {
			n, err := strconv.Atoi(col)
			if err != nil {
				fmt.Println("無効な表現です。:", col)
				os.Exit(1)
			}
			if (n < 0) || (n > len(head)-1) {
				fmt.Println("無効なカラム番号です。:", n)
				os.Exit(1)
			}
			nCols = append(nCols, n-1)
		}

		// ********** orders
		bOrdrs := []bool{}
		for _, o := range orders {
			if o == "a" {
				bOrdrs = append(bOrdrs, true)
			} else if o == "d" {
				bOrdrs = append(bOrdrs, false)
			} else {
				fmt.Println("無効なオーダーです。:", o)
				os.Exit(1)
			}
		}

		cmp := sort.MakeCmp(nCols, bOrdrs)
		start_t := time.Now()
		result := sort.GoMergeSort(cmp, dat)
		sorting_time := time.Now().Sub(start_t).Seconds()
		fmt.Println("***sorting time:", sorting_time*(1000), "ms")

		// ******************** writing ********************
		write_file, err := os.OpenFile(sortedAddrBookPath, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		mycsv.Write(head, result, write_file)

		return nil
	}
	app.Run(os.Args)
}
