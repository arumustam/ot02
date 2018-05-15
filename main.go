package main

import (
  "fmt"
  "os"
  "strings"
  _"reflect"

  "./encoder"
)

type index int8
type blood int8
type gender int8

// index
const (
  id index = iota
  fullNameKanji
  lastNameKatakana
  firstNameKatakana
  genderNum
  bloodTypeNum
  birthday
  tel
  mobile
  mailAddr
  zipCode
  addrKanji
  addrKatakana
)

// blood 
const (
  aType blood = 1 + iota
  bType
  oType
  abType
)

// gender
const (
  male gender = 1 + iota
  female
)

func split(csv_rows []string) [][]string {
  var ret [][]string
  for _, row := range(csv_rows) {
    splited_row := strings.Split(row, ",")
    ret = append(ret, splited_row)
  }
  return ret
}

func main() {
  //write_file, _ := os.OpenFile("./out", os.O_WRONLY|os.O_CREATE, 0600)

  read_file, _ := os.OpenFile("./AddrBook.csv", os.O_RDONLY, 0600)
  dat := split(encoder.ToUTF8(read_file))
  fmt.Println(dat[8])
}
