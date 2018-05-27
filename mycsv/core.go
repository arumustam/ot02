package mycsv

import (
	"bufio"
	"encoding/csv"
	_ "fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	Id = iota
	FullNameKanji
	LastNameKatakana
	FirstNameKatakana
	GenderNum
	BloodTypeNum
	Birthday
	Tel
	Mobile
	MailAddr
	ZipCode
	AddrKanji
	AddrKatakana
)

func Read(inStream io.Reader) ([]string, [][]string) {
	// **** encode ****
	var dat []string
	scanner := bufio.NewScanner(transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder()))
	for scanner.Scan() {
		dat = append(dat, scanner.Text())
	}
	// **** split ****
	head := strings.Split(dat[0], ",")
	var ret [][]string
	for i := 1; i < len(dat); i++ {
		splited_row := strings.Split(dat[i], ",")
		ret = append(ret, splited_row)
	}
	return head, ret
}

func Write(head []string, dat [][]string, outStream io.Writer) {
	writer := csv.NewWriter(transform.NewWriter(outStream, japanese.ShiftJIS.NewEncoder()))
	writer.Write(head)
	for i, row := range dat {
		row[Id] = strconv.Itoa(i + 1)
		writer.Write(row)
	}
	writer.Flush()
}
